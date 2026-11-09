package httpadapter

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	graphqladapter "seller-dashboard/internal/adapters/inbound/graphql"
	"seller-dashboard/internal/application"
	"seller-dashboard/internal/domain"
)

type Router struct {
	s   *application.Service
	gql http.Handler
}

func New(s *application.Service) http.Handler {
	r := &Router{s: s, gql: graphqladapter.New(s)}
	return r.cors(http.HandlerFunc(r.route))
}
func (r *Router) cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		o := q.Header.Get("Origin")
		preview := strings.HasPrefix(o, "https://") && strings.HasSuffix(strings.TrimPrefix(o, "https://"), "-3000.us-vmprovider.projects.hrcdn.net")
		if o == "http://localhost:3000" || o == "http://127.0.0.1:3000" || preview {
			w.Header().Set("Access-Control-Allow-Origin", o)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		}
		if q.Method == http.MethodOptions {
			w.WriteHeader(204)
			return
		}
		next.ServeHTTP(w, q)
	})
}
func (r *Router) route(w http.ResponseWriter, q *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := strings.Trim(q.URL.Path, "/")
	parts := strings.Split(p, "/")
	switch {
	case q.Method == "GET" && p == "api/health":
		write(w, 200, map[string]string{"status": "ok"})
	case q.Method == "POST" && p == "api/auth/login":
		r.login(w, q)
	case q.Method == "GET" && p == "api/roles/seeded-users":
		r.users(w)
	case q.Method == "POST" && p == "graphql":
		r.gql.ServeHTTP(w, q)
	case len(parts) == 4 && parts[0] == "api" && parts[1] == "sellers" && parts[3] == "orders" && q.Method == "GET":
		r.sellerOrders(w, q, parts[2])
	case len(parts) == 5 && parts[0] == "api" && parts[1] == "sellers" && parts[3] == "orders" && q.Method == "GET":
		r.sellerOrder(w, q, parts[2], parts[4])
	case len(parts) == 6 && parts[0] == "api" && parts[1] == "sellers" && parts[3] == "questions" && parts[5] == "reply" && q.Method == "POST":
		r.reply(w, q, parts[2], parts[4])
	case len(parts) == 4 && parts[0] == "api" && parts[1] == "orders" && parts[3] == "questions" && q.Method == "POST":
		r.backendQuestion(w, q, parts[2])
	case len(parts) == 4 && parts[0] == "api" && parts[1] == "buyers" && parts[3] == "orders" && q.Method == "GET":
		r.buyerOrders(w, q, parts[2])
	case len(parts) == 5 && parts[0] == "api" && parts[1] == "buyers" && parts[3] == "orders" && q.Method == "GET":
		r.buyerOrder(w, q, parts[2], parts[4])
	case len(parts) == 6 && parts[0] == "api" && parts[1] == "buyers" && parts[3] == "orders" && parts[5] == "questions" && q.Method == "POST":
		r.buyerQuestion(w, q, parts[2], parts[4])
	case len(parts) == 6 && parts[0] == "api" && parts[1] == "buyers" && parts[3] == "questions" && parts[5] == "approve-answer" && q.Method == "POST":
		r.approve(w, q, parts[2], parts[4])
	case p == "api/operations/questions" && q.Method == "GET":
		r.operationsQuestions(w, q)
	case p == "api/operations/priority-review" && q.Method == "POST":
		r.review(w, q)
	default:
		fail(w, 404, "not_found", "Route not found")
	}
}
func decode(q *http.Request, v any) error {
	d := json.NewDecoder(q.Body)
	d.DisallowUnknownFields()
	return d.Decode(v)
}
func write(w http.ResponseWriter, status int, v any) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func fail(w http.ResponseWriter, status int, code, msg string) {
	write(w, status, map[string]any{"error": map[string]any{"code": code, "message": msg, "details": map[string]any{}}})
}
func appError(w http.ResponseWriter, e error) {
	switch {
	case errors.Is(e, application.ErrValidation):
		fail(w, 400, "validation_error", "Request validation failed")
	case errors.Is(e, application.ErrForbidden):
		fail(w, 403, "forbidden", "Access denied")
	case errors.Is(e, application.ErrNotFound):
		fail(w, 404, "not_found", "Resource not found")
	case errors.Is(e, application.ErrConflict):
		fail(w, 409, "conflict", "Action conflicts with current lifecycle")
	default:
		fail(w, 500, "internal_error", "Unexpected error")
	}
}
func (r *Router) auth(q *http.Request) (domain.Account, error) {
	header := q.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		return domain.Account{}, application.ErrForbidden
	}
	token := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
	if token == "" {
		return domain.Account{}, application.ErrForbidden
	}
	return r.s.Authenticate(token)
}
func (r *Router) login(w http.ResponseWriter, q *http.Request) {
	var x struct{ Email, Password string }
	if decode(q, &x) != nil {
		fail(w, 400, "bad_request", "Malformed request")
		return
	}
	a, t, e := r.s.Login(x.Email, x.Password)
	if e != nil {
		fail(w, 401, "invalid_credentials", "Invalid email or password")
		return
	}
	write(w, 200, map[string]any{"user": accountJSON(a), "accessToken": t})
}
func accountJSON(a domain.Account) map[string]any {
	return map[string]any{"id": a.ID, "email": a.Email, "displayName": a.DisplayName, "roles": a.Roles, "sellerId": emptyNil(a.SellerID), "opsTeamMemberId": emptyNil(a.OpsTeamMemberID), "buyerId": emptyNil(a.BuyerID)}
}
func emptyNil(s string) any {
	if s == "" {
		return nil
	}
	return s
}
func (r *Router) users(w http.ResponseWriter) {
	out := []any{}
	for _, a := range r.s.SeededUsers() {
		out = append(out, accountJSON(a))
	}
	write(w, 200, map[string]any{"users": out})
}
func filters(q *http.Request) application.Filters {
	v := q.URL.Query()
	return application.Filters{Status: v.Get("status"), Buyer: v.Get("buyer"), Seller: v.Get("seller"), DateFrom: v.Get("dateFrom"), DateTo: v.Get("dateTo"), PriorityClass: v.Get("priorityClass")}
}
func (r *Router) sellerOrders(w http.ResponseWriter, q *http.Request, id string) {
	a, e := r.auth(q)
	if e != nil {
		fail(w, 401, "unauthorized", "Authentication required")
		return
	}
	if e := r.s.AuthorizeSeller(a, id); e != nil {
		appError(w, e)
		return
	}
	o, e := r.s.SellerOrders(id, filters(q))
	if e != nil {
		appError(w, e)
		return
	}
	write(w, 200, map[string]any{"orders": ordersJSON(r.s, o)})
}
func (r *Router) sellerOrder(w http.ResponseWriter, q *http.Request, sid, oid string) {
	a, e := r.auth(q)
	if e != nil {
		fail(w, 401, "unauthorized", "Authentication required")
		return
	}
	if e := r.s.AuthorizeSeller(a, sid); e != nil {
		appError(w, e)
		return
	}
	o, e := r.s.SellerOrder(sid, oid)
	if e != nil {
		appError(w, e)
		return
	}
	write(w, 200, map[string]any{"order": orderJSON(r.s, o)})
}
func (r *Router) reply(w http.ResponseWriter, q *http.Request, sid, qid string) {
	a, e := r.auth(q)
	if e != nil {
		fail(w, 401, "unauthorized", "Authentication required")
		return
	}
	if e := r.s.AuthorizeSeller(a, sid); e != nil {
		appError(w, e)
		return
	}
	var x struct{ Body string }
	if decode(q, &x) != nil {
		fail(w, 400, "bad_request", "Malformed request")
		return
	}
	v, e := r.s.Reply(sid, qid, x.Body)
	if e != nil {
		appError(w, e)
		return
	}
	write(w, 201, map[string]any{"question": questionJSON(v)})
}
func (r *Router) backendQuestion(w http.ResponseWriter, q *http.Request, oid string) {
	var x struct{ BuyerID, Body, OrderItemID string }
	if decode(q, &x) != nil {
		fail(w, 400, "bad_request", "Malformed request")
		return
	}
	v, e := r.s.CreateQuestion(nil, oid, x.BuyerID, x.Body, x.OrderItemID)
	if e != nil {
		appError(w, e)
		return
	}
	write(w, 201, map[string]any{"question": questionJSON(v), "notification": map[string]any{"eligible": v.PriorityClass == "High Priority" || v.PriorityClass == "Critical"}})
}
func (r *Router) buyerOrders(w http.ResponseWriter, q *http.Request, id string) {
	a, e := r.auth(q)
	if e != nil {
		fail(w, 401, "unauthorized", "Authentication required")
		return
	}
	o, e := r.s.BuyerOrders(a, id, filters(q))
	if e != nil {
		appError(w, e)
		return
	}
	write(w, 200, map[string]any{"orders": ordersJSON(r.s, o)})
}
func (r *Router) buyerOrder(w http.ResponseWriter, q *http.Request, bid, oid string) {
	a, e := r.auth(q)
	if e != nil {
		fail(w, 401, "unauthorized", "Authentication required")
		return
	}
	o, e := r.s.BuyerOrder(a, bid, oid)
	if e != nil {
		appError(w, e)
		return
	}
	write(w, 200, map[string]any{"order": orderJSON(r.s, o)})
}
func (r *Router) buyerQuestion(w http.ResponseWriter, q *http.Request, bid, oid string) {
	a, e := r.auth(q)
	if e != nil {
		fail(w, 401, "unauthorized", "Authentication required")
		return
	}
	var x struct{ Body, OrderItemID string }
	if decode(q, &x) != nil {
		fail(w, 400, "bad_request", "Malformed request")
		return
	}
	v, e := r.s.CreateQuestion(&a, oid, bid, x.Body, x.OrderItemID)
	if e != nil {
		appError(w, e)
		return
	}
	write(w, 201, map[string]any{"question": questionJSON(v)})
}
func (r *Router) approve(w http.ResponseWriter, q *http.Request, bid, qid string) {
	a, e := r.auth(q)
	if e != nil {
		fail(w, 401, "unauthorized", "Authentication required")
		return
	}
	v, e := r.s.Approve(a, bid, qid)
	if e != nil {
		appError(w, e)
		return
	}
	write(w, 200, map[string]any{"question": questionJSON(v)})
}
func (r *Router) review(w http.ResponseWriter, q *http.Request) {
	a, e := r.auth(q)
	if e != nil {
		fail(w, 401, "unauthorized", "Authentication required")
		return
	}
	if r.s.AuthorizeOperations(a) != nil {
		fail(w, 403, "forbidden", "Operations role required")
		return
	}
	var x struct{ ReviewedAt string }
	if e := decode(q, &x); e != nil && q.ContentLength != 0 {
		fail(w, 400, "bad_request", "Malformed request")
		return
	}
	at := time.Now().UTC()
	if x.ReviewedAt != "" {
		v, e := time.Parse(time.RFC3339, x.ReviewedAt)
		if e != nil {
			fail(w, 400, "validation_error", "reviewedAt must be an ISO 8601 timestamp")
			return
		}
		at = v
	}
	result := r.s.ReviewManual(at)
	write(w, 200, map[string]any{"review": resultJSON(result)})
}

func (r *Router) operationsQuestions(w http.ResponseWriter, q *http.Request) {
	a, e := r.auth(q)
	if e != nil {
		fail(w, 401, "unauthorized", "Authentication required")
		return
	}
	if e := r.s.AuthorizeOperations(a); e != nil {
		appError(w, e)
		return
	}
	write(w, 200, map[string]any{"questions": queueJSON(r.s.Queue(filters(q)))})
}

func resultJSON(v application.ReviewResult) map[string]any {
	return map[string]any{
		"reviewedAt": v.ReviewedAt, "reviewWindow": v.ReviewWindow, "executionId": v.ExecutionID,
		"status": v.Status, "eligibleOpenQuestions": v.EligibleOpenQuestions,
		"recategorized": v.Recategorized, "notificationBatchesLogged": v.NotificationBatchesLogged,
	}
}
func ordersJSON(s *application.Service, os []domain.Order) []any {
	out := []any{}
	for _, o := range os {
		out = append(out, orderJSON(s, o))
	}
	return out
}
func orderJSON(s *application.Service, o domain.Order) map[string]any {
	items := []any{}
	for _, i := range o.Items {
		items = append(items, map[string]any{"id": i.ID, "productId": i.ProductID, "productName": i.ProductName, "quantity": i.Quantity, "unitPrice": i.UnitPrice})
	}
	qs := []any{}
	for _, q := range s.QuestionsForOrder(o.ID) {
		qs = append(qs, questionJSON(q))
	}
	return map[string]any{
		"id": o.ID, "sellerId": o.SellerID, "buyerId": o.BuyerID,
		"seller": s.Seller(o.SellerID), "buyer": s.Buyer(o.BuyerID),
		"status": o.Status, "orderedAt": o.OrderedAt, "items": items,
		"total": o.Total(), "questions": qs,
	}
}
func questionJSON(q domain.Question) map[string]any {
	return map[string]any{"id": q.ID, "orderId": q.OrderID, "buyerId": q.BuyerID, "orderItemId": emptyNil(q.OrderItemID), "body": q.Body, "status": q.Status, "lifecycleStatus": q.LifecycleStatus, "createdAt": q.CreatedAt, "buyerApprovedAt": q.BuyerApprovedAt, "resolvedAt": q.ResolvedAt, "reply": q.Reply, "priorityScore": q.PriorityScore, "priorityClass": q.PriorityClass, "matchedUrgencyWords": q.MatchedUrgencyWords, "priorityDictionaryVersion": q.PriorityDictionaryVersion, "lastCategorizedAt": q.LastCategorizedAt}
}
func queueJSON(xs []application.QueueItem) []any {
	out := []any{}
	for _, x := range xs {
		out = append(out, map[string]any{"question": questionJSON(x.Question), "order": map[string]any{"id": x.Order.ID, "status": x.Order.Status, "orderedAt": x.Order.OrderedAt, "total": x.Order.Total()}, "seller": x.Seller, "buyer": x.Buyer, "waitingDays": x.WaitingDays})
	}
	return out
}
