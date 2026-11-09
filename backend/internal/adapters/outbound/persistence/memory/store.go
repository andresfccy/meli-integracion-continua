package memory

import (
	"sync"
	"time"

	"seller-dashboard/internal/domain"
)

type Store struct {
	mu            sync.RWMutex
	sellers       []domain.Seller
	buyers        []domain.Buyer
	products      []domain.Product
	opsMembers    []domain.OpsTeamMember
	notifiables   []domain.Notifiable
	orders        []domain.Order
	questions     []domain.Question
	accounts      []domain.Account
	channels      []domain.NotificationChannel
	events        []domain.NotificationEvent
	reviewWindows map[string]struct{}
}

func New() *Store {
	now := time.Now().UTC()
	s := &Store{reviewWindows: map[string]struct{}{}}
	s.sellers = []domain.Seller{{ID: "seller-1", Name: "Mercado Norte", Email: "seller@example.com", NotificationChannelIDs: []string{"channel-1"}}, {ID: "seller-2", Name: "Tienda Andina", Email: "seller2@example.com", NotificationChannelIDs: []string{"channel-2"}}}
	s.buyers = []domain.Buyer{{ID: "buyer-1", Name: "Ana García", Email: "buyer@example.com"}, {ID: "buyer-2", Name: "Luis Pérez", Email: "buyer2@example.com"}}
	s.products = []domain.Product{
		{ID: "product-1", SellerID: "seller-1", Name: "Auriculares inalámbricos", Category: "audio", CurrentPrice: 139.90},
		{ID: "product-2", SellerID: "seller-1", Name: "Funda protectora", Category: "accessories", CurrentPrice: 24.50},
		{ID: "product-3", SellerID: "seller-1", Name: "Teclado mecánico", Category: "computing", CurrentPrice: 269.00},
		{ID: "product-4", SellerID: "seller-1", Name: "Monitor 27 pulgadas", Category: "computing", CurrentPrice: 949.00},
		{ID: "product-5", SellerID: "seller-2", Name: "Cafetera", Category: "home", CurrentPrice: 350.00},
	}
	s.opsMembers = []domain.OpsTeamMember{{ID: "ops-1", Name: "Operaciones", Email: "ops@example.com", NotificationChannelIDs: []string{"channel-ops-1"}}}
	s.accounts = []domain.Account{{ID: "user-seller-1", Email: "seller@example.com", Password: "seller-demo-password", DisplayName: "Mercado Norte", Roles: []string{"seller"}, SellerID: "seller-1"}, {ID: "user-seller-2", Email: "seller2@example.com", Password: "seller2-demo-password", DisplayName: "Tienda Andina", Roles: []string{"seller"}, SellerID: "seller-2"}, {ID: "user-ops", Email: "ops@example.com", Password: "ops-demo-password", DisplayName: "Operaciones", Roles: []string{"operations"}, OpsTeamMemberID: "ops-1"}, {ID: "user-buyer", Email: "buyer@example.com", Password: "buyer-demo-password", DisplayName: "Ana García", Roles: []string{"buyer"}, BuyerID: "buyer-1"}}
	s.channels = []domain.NotificationChannel{{ID: "channel-1", Type: "email", Destination: "seller@example.com", Enabled: true, Label: "Console-simulated primary email"}, {ID: "channel-2", Type: "email", Destination: "seller2@example.com", Enabled: true, Label: "Console-simulated primary email"}, {ID: "channel-ops-1", Type: "email", Destination: "ops@example.com", Enabled: true, Label: "Console-simulated Operations email"}}
	factory := domain.NewNotifiableFactory()
	for _, seller := range s.sellers {
		target, err := factory.FromSeller(seller)
		if err != nil {
			panic(err)
		}
		s.notifiables = append(s.notifiables, target)
	}
	for _, member := range s.opsMembers {
		target, err := factory.FromOpsTeamMember(member)
		if err != nil {
			panic(err)
		}
		s.notifiables = append(s.notifiables, target)
	}
	s.orders = []domain.Order{{ID: "order-1001", SellerID: "seller-1", BuyerID: "buyer-1", Status: "processing", OrderedAt: now.Add(-72 * time.Hour), Items: []domain.OrderItem{{ID: "item-1001", OrderID: "order-1001", ProductID: "product-1", ProductName: "Auriculares inalámbricos", Quantity: 1, UnitPrice: 129.90}, {ID: "item-1002", OrderID: "order-1001", ProductID: "product-2", ProductName: "Funda protectora", Quantity: 2, UnitPrice: 19.50}}}, {ID: "order-1002", SellerID: "seller-1", BuyerID: "buyer-2", Status: "shipped", OrderedAt: now.Add(-24 * time.Hour), Items: []domain.OrderItem{{ID: "item-1003", OrderID: "order-1002", ProductID: "product-3", ProductName: "Teclado mecánico", Quantity: 1, UnitPrice: 249.00}}}, {ID: "order-1003", SellerID: "seller-1", BuyerID: "buyer-1", Status: "delivered", OrderedAt: now.Add(-240 * time.Hour), Items: []domain.OrderItem{{ID: "item-1004", OrderID: "order-1003", ProductID: "product-4", ProductName: "Monitor 27 pulgadas", Quantity: 1, UnitPrice: 899.00}}}, {ID: "order-2001", SellerID: "seller-2", BuyerID: "buyer-2", Status: "new", OrderedAt: now.Add(-12 * time.Hour), Items: []domain.OrderItem{{ID: "item-2001", OrderID: "order-2001", ProductID: "product-5", ProductName: "Cafetera", Quantity: 1, UnitPrice: 320}}}}
	score, class, words := domain.EvaluatePriority("Urgent refund: my package never arrived", now.Add(-72*time.Hour), now, 168.90)
	resolvedAt := now.Add(-12 * time.Hour)
	s.questions = []domain.Question{
		{ID: "question-1", OrderID: "order-1001", BuyerID: "buyer-1", OrderItemID: "item-1001", Body: "Urgent refund: my package never arrived", Status: "open", LifecycleStatus: "question_active", CreatedAt: now.Add(-72 * time.Hour), PriorityScore: score, PriorityClass: class, MatchedUrgencyWords: words, PriorityDictionaryVersion: "mvp-2026-07-13", LastCategorizedAt: now.Add(-72 * time.Hour)},
		{ID: "question-2", OrderID: "order-1002", BuyerID: "buyer-2", Body: "Question already resolved", Status: "resolved", LifecycleStatus: "closed", CreatedAt: now.Add(-48 * time.Hour), BuyerApprovedAt: &resolvedAt, ResolvedAt: &resolvedAt, Reply: &domain.Reply{ID: "reply-question-2", QuestionID: "question-2", SellerID: "seller-1", Body: "Resolved answer", CreatedAt: now.Add(-24 * time.Hour)}, PriorityScore: 0, PriorityClass: "Low", MatchedUrgencyWords: []string{}, PriorityDictionaryVersion: "mvp-2026-07-13", LastCategorizedAt: now.Add(-48 * time.Hour)},
	}
	s.events = []domain.NotificationEvent{{ID: "event-1", QuestionIDs: []string{"question-2"}, NotifiableID: "notifiable-seller-seller-1", ChannelID: "channel-1", BatchDate: resolvedAt.Format("2006-01-02"), GroupingKey: "seeded:resolved", IncludedPriorityClasses: []string{"High Priority"}, Message: "Seeded historical notification event", CreatedAt: resolvedAt}}
	return s
}

func (s *Store) Sellers() []domain.Seller {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.Seller, len(s.sellers))
	for index, seller := range s.sellers {
		out[index] = seller
		out[index].NotificationChannelIDs = append([]string(nil), seller.NotificationChannelIDs...)
	}
	return out
}
func (s *Store) Buyers() []domain.Buyer {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Buyer(nil), s.buyers...)
}
func (s *Store) Products() []domain.Product {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.Product(nil), s.products...)
}
func (s *Store) OpsTeamMembers() []domain.OpsTeamMember {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.OpsTeamMember, len(s.opsMembers))
	for index, member := range s.opsMembers {
		out[index] = member
		out[index].NotificationChannelIDs = append([]string(nil), member.NotificationChannelIDs...)
	}
	return out
}
func (s *Store) Notifiables() []domain.Notifiable {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.Notifiable, len(s.notifiables))
	for index, target := range s.notifiables {
		out[index] = target
		out[index].NotificationChannelIDs = append([]string(nil), target.NotificationChannelIDs...)
	}
	return out
}
func (s *Store) Orders() []domain.Order {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.Order, len(s.orders))
	for index, order := range s.orders {
		out[index] = order
		out[index].Items = append([]domain.OrderItem(nil), order.Items...)
	}
	return out
}
func (s *Store) Questions() []domain.Question {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.Question, len(s.questions))
	for index, question := range s.questions {
		out[index] = cloneQuestion(question)
	}
	return out
}
func (s *Store) Accounts() []domain.Account {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.Account, len(s.accounts))
	for index, account := range s.accounts {
		out[index] = account
		out[index].Roles = append([]string(nil), account.Roles...)
	}
	return out
}
func (s *Store) Channels() []domain.NotificationChannel {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]domain.NotificationChannel(nil), s.channels...)
}
func (s *Store) Events() []domain.NotificationEvent {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]domain.NotificationEvent, len(s.events))
	for index, event := range s.events {
		out[index] = event
		out[index].QuestionIDs = append([]string(nil), event.QuestionIDs...)
		out[index].IncludedPriorityClasses = append([]string(nil), event.IncludedPriorityClasses...)
	}
	return out
}
func (s *Store) SaveQuestion(q domain.Question) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, x := range s.questions {
		if x.ID == q.ID {
			s.questions[i] = cloneQuestion(q)
			return
		}
	}
	s.questions = append(s.questions, cloneQuestion(q))
}
func (s *Store) SaveEvent(e domain.NotificationEvent) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e.QuestionIDs = append([]string(nil), e.QuestionIDs...)
	e.IncludedPriorityClasses = append([]string(nil), e.IncludedPriorityClasses...)
	s.events = append(s.events, e)
}

func cloneQuestion(question domain.Question) domain.Question {
	question.MatchedUrgencyWords = append([]string(nil), question.MatchedUrgencyWords...)
	if question.Reply != nil {
		reply := *question.Reply
		question.Reply = &reply
	}
	if question.BuyerApprovedAt != nil {
		approved := *question.BuyerApprovedAt
		question.BuyerApprovedAt = &approved
	}
	if question.ResolvedAt != nil {
		resolved := *question.ResolvedAt
		question.ResolvedAt = &resolved
	}
	return question
}

func (s *Store) AcquireReviewWindow(window string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.reviewWindows[window]; exists {
		return false
	}
	s.reviewWindows[window] = struct{}{}
	return true
}
