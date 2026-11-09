package application

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"seller-dashboard/internal/domain"
)

var ErrNotFound = errors.New("not found")
var ErrForbidden = errors.New("forbidden")
var ErrConflict = errors.New("conflict")
var ErrValidation = errors.New("validation")

type Store interface {
	Sellers() []domain.Seller
	Buyers() []domain.Buyer
	Products() []domain.Product
	OpsTeamMembers() []domain.OpsTeamMember
	Notifiables() []domain.Notifiable
	Orders() []domain.Order
	Questions() []domain.Question
	Accounts() []domain.Account
	Channels() []domain.NotificationChannel
	SaveQuestion(domain.Question)
	SaveEvent(domain.NotificationEvent)
	Events() []domain.NotificationEvent
	AcquireReviewWindow(string) bool
}
type Delivery interface {
	Deliver(domain.NotificationEvent, domain.NotificationChannel, string) error
}
type Service struct {
	store    Store
	delivery Delivery
	secret   []byte
	now      func() time.Time
}

func New(store Store, d Delivery) *Service {
	return &Service{store: store, delivery: d, secret: []byte("seeded-mvp-signing-secret"), now: time.Now}
}
func (s *Service) SetNow(f func() time.Time)     { s.now = f }
func (s *Service) SeededUsers() []domain.Account { return s.store.Accounts() }
func (s *Service) Login(email, password string) (domain.Account, string, error) {
	for _, a := range s.store.Accounts() {
		if strings.EqualFold(a.Email, strings.TrimSpace(email)) && hmac.Equal([]byte(a.Password), []byte(password)) {
			return a, s.token(a), nil
		}
	}
	return domain.Account{}, "", ErrForbidden
}
func (s *Service) token(a domain.Account) string {
	headerJSON, _ := json.Marshal(map[string]string{"alg": "HS256", "typ": "JWT"})
	header := b64(headerJSON)
	claims, _ := json.Marshal(map[string]any{"sub": a.ID, "iss": "seller-dashboard-mvp", "aud": "graphql-read", "iat": s.now().Unix(), "exp": s.now().Add(time.Hour).Unix()})
	payload := b64(claims)
	mac := hmac.New(sha256.New, s.secret)
	mac.Write([]byte(header + "." + payload))
	return header + "." + payload + "." + b64(mac.Sum(nil))
}
func b64(v []byte) string { return base64.RawURLEncoding.EncodeToString(v) }
func (s *Service) Authenticate(token string) (domain.Account, error) {
	p := strings.Split(token, ".")
	if len(p) != 3 {
		return domain.Account{}, ErrForbidden
	}
	m := hmac.New(sha256.New, s.secret)
	m.Write([]byte(p[0] + "." + p[1]))
	sig, e := base64.RawURLEncoding.DecodeString(p[2])
	if e != nil || !hmac.Equal(sig, m.Sum(nil)) {
		return domain.Account{}, ErrForbidden
	}
	raw, e := base64.RawURLEncoding.DecodeString(p[1])
	if e != nil {
		return domain.Account{}, ErrForbidden
	}
	var c struct {
		Sub, Iss, Aud string
		Iat, Exp      int64
	}
	headerRaw, e := base64.RawURLEncoding.DecodeString(p[0])
	var header struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}
	if e != nil || json.Unmarshal(headerRaw, &header) != nil || header.Alg != "HS256" || header.Typ != "JWT" ||
		json.Unmarshal(raw, &c) != nil || c.Iss != "seller-dashboard-mvp" || c.Aud != "graphql-read" ||
		c.Iat <= 0 || c.Iat > s.now().Unix() || c.Exp <= s.now().Unix() || c.Exp <= c.Iat {
		return domain.Account{}, ErrForbidden
	}
	for _, a := range s.store.Accounts() {
		if a.ID == c.Sub {
			return a, nil
		}
	}
	return domain.Account{}, ErrForbidden
}
func hasRole(a domain.Account, r string) bool {
	for _, x := range a.Roles {
		if x == r {
			return true
		}
	}
	return false
}

func (s *Service) AuthorizeSeller(a domain.Account, sellerID string) error {
	if !hasRole(a, "seller") || a.SellerID != sellerID {
		return ErrForbidden
	}
	return nil
}

func (s *Service) AuthorizeOperations(a domain.Account) error {
	if !hasRole(a, "operations") {
		return ErrForbidden
	}
	return nil
}

type Filters struct{ Status, Buyer, Seller, DateFrom, DateTo, PriorityClass string }

func (s *Service) SellerOrders(id string, f Filters) ([]domain.Order, error) {
	if validateFilters(f) != nil {
		return nil, ErrValidation
	}
	found := false
	for _, v := range s.store.Sellers() {
		if v.ID == id {
			found = true
		}
	}
	if !found {
		return nil, ErrNotFound
	}
	out := []domain.Order{}
	for _, o := range s.store.Orders() {
		if o.SellerID == id && s.match(o, f) {
			out = append(out, o)
		}
	}
	sortOrders(out)
	return out, nil
}
func (s *Service) BuyerOrders(a domain.Account, id string, f Filters) ([]domain.Order, error) {
	if !hasRole(a, "buyer") || a.BuyerID != id {
		return nil, ErrForbidden
	}
	if validateFilters(f) != nil {
		return nil, ErrValidation
	}
	out := []domain.Order{}
	for _, o := range s.store.Orders() {
		if o.BuyerID == id && s.match(o, f) {
			out = append(out, o)
		}
	}
	sortOrders(out)
	return out, nil
}

func validateFilters(f Filters) error {
	if f.Status != "" {
		valid := map[string]bool{"new": true, "processing": true, "shipped": true, "delivered": true, "cancelled": true}
		if !valid[f.Status] {
			return ErrValidation
		}
	}
	var from, to time.Time
	var err error
	if f.DateFrom != "" {
		from, err = time.Parse("2006-01-02", f.DateFrom)
		if err != nil {
			return ErrValidation
		}
	}
	if f.DateTo != "" {
		to, err = time.Parse("2006-01-02", f.DateTo)
		if err != nil {
			return ErrValidation
		}
	}
	if !from.IsZero() && !to.IsZero() && from.After(to) {
		return ErrValidation
	}
	return nil
}
func sortOrders(v []domain.Order) {
	sort.Slice(v, func(i, j int) bool {
		if v[i].OrderedAt.Equal(v[j].OrderedAt) {
			return v[i].ID < v[j].ID
		}
		return v[i].OrderedAt.After(v[j].OrderedAt)
	})
}
func (s *Service) match(o domain.Order, f Filters) bool {
	if f.Status != "" && o.Status != f.Status {
		return false
	}
	b := s.buyer(o.BuyerID)
	if f.Buyer != "" && !strings.Contains(strings.ToLower(b.Name+" "+b.Email), strings.ToLower(f.Buyer)) {
		return false
	}
	sel := s.seller(o.SellerID)
	if f.Seller != "" && !strings.Contains(strings.ToLower(sel.ID+" "+sel.Name+" "+sel.Email), strings.ToLower(f.Seller)) {
		return false
	}
	if f.DateFrom != "" && o.OrderedAt.Format("2006-01-02") < f.DateFrom {
		return false
	}
	if f.DateTo != "" && o.OrderedAt.Format("2006-01-02") > f.DateTo {
		return false
	}
	return true
}
func (s *Service) Order(id string) (domain.Order, error) {
	for _, o := range s.store.Orders() {
		if o.ID == id {
			return o, nil
		}
	}
	return domain.Order{}, ErrNotFound
}
func (s *Service) SellerOrder(seller, id string) (domain.Order, error) {
	o, e := s.Order(id)
	if e != nil || o.SellerID != seller {
		return domain.Order{}, ErrNotFound
	}
	return o, nil
}
func (s *Service) BuyerOrder(a domain.Account, buyer, id string) (domain.Order, error) {
	if !hasRole(a, "buyer") || a.BuyerID != buyer {
		return domain.Order{}, ErrForbidden
	}
	o, e := s.Order(id)
	if e != nil {
		return o, e
	}
	if o.BuyerID != buyer {
		return domain.Order{}, ErrNotFound
	}
	return o, nil
}
func (s *Service) QuestionsForOrder(id string) []domain.Question {
	out := []domain.Question{}
	for _, q := range s.store.Questions() {
		if q.OrderID == id {
			out = append(out, q)
		}
	}
	return out
}

func (s *Service) Seller(id string) domain.Seller { return s.seller(id) }
func (s *Service) Buyer(id string) domain.Buyer   { return s.buyer(id) }
func (s *Service) CreateQuestion(a *domain.Account, orderID, buyerID, body, itemID string) (domain.Question, error) {
	o, e := s.Order(orderID)
	if e != nil {
		return domain.Question{}, e
	}
	if a != nil && (!hasRole(*a, "buyer") || a.BuyerID != buyerID || o.BuyerID != buyerID) {
		return domain.Question{}, ErrForbidden
	}
	if strings.TrimSpace(body) == "" {
		return domain.Question{}, ErrValidation
	}
	if s.buyer(buyerID).ID == "" {
		return domain.Question{}, ErrNotFound
	}
	if itemID != "" {
		ok := false
		for _, i := range o.Items {
			if i.ID == itemID {
				ok = true
			}
		}
		if !ok {
			return domain.Question{}, ErrValidation
		}
	}
	for _, q := range s.store.Questions() {
		if q.OrderID == orderID && q.BuyerID == buyerID && q.Status == "open" {
			return domain.Question{}, ErrConflict
		}
	}
	now := s.now().UTC()
	score, class, words := domain.EvaluatePriority(body, now, now, o.Total())
	q := domain.Question{ID: fmt.Sprintf("question-%d", len(s.store.Questions())+1), OrderID: orderID, BuyerID: buyerID, OrderItemID: itemID, Body: strings.TrimSpace(body), Status: "open", LifecycleStatus: "question_active", CreatedAt: now, PriorityScore: score, PriorityClass: class, MatchedUrgencyWords: words, PriorityDictionaryVersion: "mvp-2026-07-13", LastCategorizedAt: now}
	s.store.SaveQuestion(q)
	return q, nil
}
func (s *Service) Reply(seller, qid, body string) (domain.Question, error) {
	if strings.TrimSpace(body) == "" {
		return domain.Question{}, ErrValidation
	}
	for _, q := range s.store.Questions() {
		if q.ID == qid {
			o, _ := s.Order(q.OrderID)
			if o.SellerID != seller {
				return domain.Question{}, ErrNotFound
			}
			if q.Status != "open" || q.Reply != nil {
				return domain.Question{}, ErrConflict
			}
			r := &domain.Reply{ID: "reply-" + qid, QuestionID: qid, SellerID: seller, Body: strings.TrimSpace(body), CreatedAt: s.now().UTC()}
			q.Reply = r
			q.LifecycleStatus = "seller_answered"
			s.store.SaveQuestion(q)
			return q, nil
		}
	}
	return domain.Question{}, ErrNotFound
}
func (s *Service) Approve(a domain.Account, buyer, qid string) (domain.Question, error) {
	if !hasRole(a, "buyer") || a.BuyerID != buyer {
		return domain.Question{}, ErrForbidden
	}
	for _, q := range s.store.Questions() {
		if q.ID == qid {
			o, _ := s.Order(q.OrderID)
			if o.BuyerID != buyer {
				return domain.Question{}, ErrForbidden
			}
			if q.Reply == nil || q.LifecycleStatus != "seller_answered" {
				return domain.Question{}, ErrConflict
			}
			now := s.now().UTC()
			q.Status = "resolved"
			q.LifecycleStatus = "closed"
			q.BuyerApprovedAt = &now
			q.ResolvedAt = &now
			s.store.SaveQuestion(q)
			return q, nil
		}
	}
	return domain.Question{}, ErrNotFound
}

type QueueItem struct {
	Question    domain.Question
	Order       domain.Order
	Seller      domain.Seller
	Buyer       domain.Buyer
	WaitingDays int
}

func (s *Service) Queue(f Filters) []QueueItem {
	out := []QueueItem{}
	for _, q := range s.store.Questions() {
		if q.Status != "open" || f.PriorityClass != "" && q.PriorityClass != f.PriorityClass {
			continue
		}
		o, _ := s.Order(q.OrderID)
		if !s.match(o, f) {
			continue
		}
		out = append(out, QueueItem{q, o, s.seller(o.SellerID), s.buyer(o.BuyerID), int(s.now().Sub(q.CreatedAt).Hours() / 24)})
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.Question.PriorityScore != b.Question.PriorityScore {
			return a.Question.PriorityScore > b.Question.PriorityScore
		}
		if !a.Question.CreatedAt.Equal(b.Question.CreatedAt) {
			return a.Question.CreatedAt.Before(b.Question.CreatedAt)
		}
		if a.Order.Total() != b.Order.Total() {
			return a.Order.Total() > b.Order.Total()
		}
		return a.Question.ID < b.Question.ID
	})
	return out
}

type ReviewResult struct {
	ReviewedAt                time.Time
	ReviewWindow              string
	ExecutionID               string
	Status                    string
	EligibleOpenQuestions     int
	Recategorized             int
	NotificationBatchesLogged int
}

func (s *Service) ReviewManual(at time.Time) ReviewResult {
	return s.review(at.UTC(), "manual", true)
}

func (s *Service) ReviewScheduled(at time.Time) ReviewResult {
	at = at.UTC()
	window := at.Format("2006-01-02")
	if !s.store.AcquireReviewWindow(window) {
		return ReviewResult{ReviewedAt: at, ReviewWindow: window, ExecutionID: "scheduled-" + window, Status: "already_executed"}
	}
	return s.review(at, "scheduled", false)
}

func (s *Service) review(at time.Time, trigger string, manualOverride bool) ReviewResult {
	result := ReviewResult{
		ReviewedAt:   at,
		ReviewWindow: at.Format("2006-01-02"),
		ExecutionID:  trigger + "-" + at.Format("20060102T150405"),
		Status:       "completed",
	}
	for _, q := range s.store.Questions() {
		if q.Status != "open" {
			continue
		}
		result.EligibleOpenQuestions++
		if at.Sub(q.LastCategorizedAt) >= 48*time.Hour {
			o, _ := s.Order(q.OrderID)
			q.PriorityScore, q.PriorityClass, q.MatchedUrgencyWords = domain.EvaluatePriority(q.Body, q.CreatedAt, at, o.Total())
			q.LastCategorizedAt = at
			s.store.SaveQuestion(q)
			result.Recategorized++
		}
	}
	groups := map[string][]QueueItem{}
	for _, x := range s.Queue(Filters{}) {
		if x.Question.PriorityClass == "High Priority" || x.Question.PriorityClass == "Critical" {
			groups[x.Seller.ID] = append(groups[x.Seller.ID], x)
		}
	}
	for sid, items := range groups {
		target := s.notifiable("seller", sid)
		if target.ID == "" {
			continue
		}
		for _, channelID := range target.NotificationChannelIDs {
			ch := s.channel(channelID)
			if ch.ID == "" || !ch.Enabled {
				continue
			}
			ids := []string{}
			classes := []string{}
			for _, x := range items {
				ids = append(ids, x.Question.ID)
				classes = append(classes, x.Question.PriorityClass)
			}
			groupingKey := target.ID + ":" + ch.ID + ":" + at.Format("2006-01-02")
			if manualOverride {
				groupingKey += fmt.Sprintf(":manual-%d", at.UnixNano())
			}
			ev := domain.NotificationEvent{ID: fmt.Sprintf("event-%d", len(s.store.Events())+1), QuestionIDs: ids, NotifiableID: target.ID, ChannelID: ch.ID, BatchDate: at.Format("2006-01-02"), GroupingKey: groupingKey, IncludedPriorityClasses: classes, Message: fmt.Sprintf("%d priority questions need attention", len(ids)), CreatedAt: at}
			s.store.SaveEvent(ev)
			if s.delivery.Deliver(ev, ch, trigger) == nil {
				result.NotificationBatchesLogged++
			}
		}
	}
	return result
}
func (s *Service) notifiable(objectType, objectID string) domain.Notifiable {
	for _, target := range s.store.Notifiables() {
		if target.ObjectType == objectType && target.ObjectID == objectID {
			return target
		}
	}
	return domain.Notifiable{}
}
func (s *Service) channel(id string) domain.NotificationChannel {
	for _, channel := range s.store.Channels() {
		if channel.ID == id {
			return channel
		}
	}
	return domain.NotificationChannel{}
}
func (s *Service) seller(id string) domain.Seller {
	for _, x := range s.store.Sellers() {
		if x.ID == id {
			return x
		}
	}
	return domain.Seller{}
}
func (s *Service) buyer(id string) domain.Buyer {
	for _, x := range s.store.Buyers() {
		if x.ID == id {
			return x
		}
	}
	return domain.Buyer{}
}
