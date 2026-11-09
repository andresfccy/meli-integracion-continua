package graphqladapter

import (
	"context"
	_ "embed"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"seller-dashboard/internal/application"
	"seller-dashboard/internal/domain"
)

//go:embed schema.graphql
var schemaSource string

type accountContextKey struct{}

type Handler struct {
	service *application.Service
	relay   *relay.Handler
}

func New(service *application.Service) *Handler {
	schema := graphql.MustParseSchema(
		schemaSource,
		&queryResolver{service: service},
		graphql.MaxDepth(20),
		graphql.MaxQueryLength(10_000),
	)
	return &Handler{service: service, relay: &relay.Handler{Schema: schema}}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	account, err := h.authenticate(request)
	if err != nil {
		writeUnauthorized(w)
		return
	}
	ctx := context.WithValue(request.Context(), accountContextKey{}, account)
	h.relay.ServeHTTP(w, request.WithContext(ctx))
}

func (h *Handler) authenticate(request *http.Request) (domain.Account, error) {
	header := request.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		return domain.Account{}, application.ErrForbidden
	}
	return h.service.Authenticate(strings.TrimPrefix(header, "Bearer "))
}

func writeUnauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"error": map[string]any{
			"code": "unauthorized", "message": "Authentication required", "details": map[string]any{},
		},
	})
}

func accountFromContext(ctx context.Context) (domain.Account, error) {
	account, ok := ctx.Value(accountContextKey{}).(domain.Account)
	if !ok {
		return domain.Account{}, application.ErrForbidden
	}
	return account, nil
}

type queryResolver struct {
	service *application.Service
}

func (r *queryResolver) Viewer(ctx context.Context) (*accountResolver, error) {
	account, err := accountFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return &accountResolver{account: account}, nil
}

func (r *queryResolver) SellerDashboard(ctx context.Context, args struct {
	SellerID string
	Filters  *filterInput
}) (*sellerDashboardResolver, error) {
	account, err := accountFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := r.service.AuthorizeSeller(account, args.SellerID); err != nil {
		return nil, err
	}
	orders, err := r.service.SellerOrders(args.SellerID, args.Filters.applicationFilters())
	if err != nil {
		return nil, err
	}
	return &sellerDashboardResolver{orders: orderResolvers(r.service, orders)}, nil
}

func (r *queryResolver) SellerOrderDetail(ctx context.Context, args struct {
	SellerID string
	OrderID  string
}) (*orderResolver, error) {
	account, err := accountFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := r.service.AuthorizeSeller(account, args.SellerID); err != nil {
		return nil, err
	}
	order, err := r.service.SellerOrder(args.SellerID, args.OrderID)
	if err != nil {
		return nil, err
	}
	return &orderResolver{service: r.service, order: order}, nil
}

func (r *queryResolver) BuyerDashboard(ctx context.Context, args struct {
	BuyerID string
	Filters *filterInput
}) (*buyerDashboardResolver, error) {
	account, err := accountFromContext(ctx)
	if err != nil {
		return nil, err
	}
	orders, err := r.service.BuyerOrders(account, args.BuyerID, args.Filters.applicationFilters())
	if err != nil {
		return nil, err
	}
	return &buyerDashboardResolver{orders: orderResolvers(r.service, orders)}, nil
}

func (r *queryResolver) BuyerOrderDetail(ctx context.Context, args struct {
	BuyerID string
	OrderID string
}) (*orderResolver, error) {
	account, err := accountFromContext(ctx)
	if err != nil {
		return nil, err
	}
	order, err := r.service.BuyerOrder(account, args.BuyerID, args.OrderID)
	if err != nil {
		return nil, err
	}
	return &orderResolver{service: r.service, order: order}, nil
}

func (r *queryResolver) OperationsPriorityQueue(ctx context.Context, args struct {
	Filters *filterInput
}) (*operationsQueueResolver, error) {
	account, err := accountFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := r.service.AuthorizeOperations(account); err != nil {
		return nil, err
	}
	items := r.service.Queue(args.Filters.applicationFilters())
	resolvers := make([]*priorityQueueItemResolver, 0, len(items))
	for _, item := range items {
		resolvers = append(resolvers, &priorityQueueItemResolver{item: item})
	}
	return &operationsQueueResolver{questions: resolvers}, nil
}

type filterInput struct {
	Status        *string
	Buyer         *string
	Seller        *string
	DateFrom      *string
	DateTo        *string
	PriorityClass *string
}

func (f *filterInput) applicationFilters() application.Filters {
	if f == nil {
		return application.Filters{}
	}
	return application.Filters{
		Status: value(f.Status), Buyer: value(f.Buyer), Seller: value(f.Seller),
		DateFrom: value(f.DateFrom), DateTo: value(f.DateTo), PriorityClass: value(f.PriorityClass),
	}
}

func value(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

type accountResolver struct{ account domain.Account }

func (r *accountResolver) ID() string               { return r.account.ID }
func (r *accountResolver) Email() string            { return r.account.Email }
func (r *accountResolver) DisplayName() string      { return r.account.DisplayName }
func (r *accountResolver) Roles() []string          { return r.account.Roles }
func (r *accountResolver) SellerID() *string        { return nullable(r.account.SellerID) }
func (r *accountResolver) BuyerID() *string         { return nullable(r.account.BuyerID) }
func (r *accountResolver) OpsTeamMemberID() *string { return nullable(r.account.OpsTeamMemberID) }

type sellerDashboardResolver struct{ orders []*orderResolver }

func (r *sellerDashboardResolver) Orders() []*orderResolver { return r.orders }

type buyerDashboardResolver struct{ orders []*orderResolver }

func (r *buyerDashboardResolver) Orders() []*orderResolver { return r.orders }

type operationsQueueResolver struct{ questions []*priorityQueueItemResolver }

func (r *operationsQueueResolver) Questions() []*priorityQueueItemResolver { return r.questions }

type orderResolver struct {
	service *application.Service
	order   domain.Order
}

func (r *orderResolver) ID() string        { return r.order.ID }
func (r *orderResolver) SellerID() string  { return r.order.SellerID }
func (r *orderResolver) BuyerID() string   { return r.order.BuyerID }
func (r *orderResolver) Status() string    { return r.order.Status }
func (r *orderResolver) OrderedAt() string { return formatTime(r.order.OrderedAt) }
func (r *orderResolver) Total() float64    { return r.order.Total() }
func (r *orderResolver) Seller() *sellerResolver {
	return &sellerResolver{seller: r.service.Seller(r.order.SellerID)}
}
func (r *orderResolver) Buyer() *buyerResolver {
	return &buyerResolver{buyer: r.service.Buyer(r.order.BuyerID)}
}
func (r *orderResolver) Items() []*orderItemResolver {
	items := make([]*orderItemResolver, 0, len(r.order.Items))
	for _, item := range r.order.Items {
		items = append(items, &orderItemResolver{item: item})
	}
	return items
}
func (r *orderResolver) Questions() []*questionResolver {
	questions := r.service.QuestionsForOrder(r.order.ID)
	resolvers := make([]*questionResolver, 0, len(questions))
	for _, question := range questions {
		resolvers = append(resolvers, &questionResolver{question: question})
	}
	return resolvers
}

func orderResolvers(service *application.Service, orders []domain.Order) []*orderResolver {
	resolvers := make([]*orderResolver, 0, len(orders))
	for _, order := range orders {
		resolvers = append(resolvers, &orderResolver{service: service, order: order})
	}
	return resolvers
}

type orderItemResolver struct{ item domain.OrderItem }

func (r *orderItemResolver) ID() string          { return r.item.ID }
func (r *orderItemResolver) ProductID() string   { return r.item.ProductID }
func (r *orderItemResolver) ProductName() string { return r.item.ProductName }
func (r *orderItemResolver) Quantity() int32     { return int32(r.item.Quantity) }
func (r *orderItemResolver) UnitPrice() float64  { return r.item.UnitPrice }

type questionResolver struct{ question domain.Question }

func (r *questionResolver) ID() string                    { return r.question.ID }
func (r *questionResolver) OrderID() string               { return r.question.OrderID }
func (r *questionResolver) BuyerID() string               { return r.question.BuyerID }
func (r *questionResolver) OrderItemID() *string          { return nullable(r.question.OrderItemID) }
func (r *questionResolver) Body() string                  { return r.question.Body }
func (r *questionResolver) Status() string                { return r.question.Status }
func (r *questionResolver) LifecycleStatus() string       { return r.question.LifecycleStatus }
func (r *questionResolver) CreatedAt() string             { return formatTime(r.question.CreatedAt) }
func (r *questionResolver) BuyerApprovedAt() *string      { return nullableTime(r.question.BuyerApprovedAt) }
func (r *questionResolver) ResolvedAt() *string           { return nullableTime(r.question.ResolvedAt) }
func (r *questionResolver) PriorityScore() int32          { return int32(r.question.PriorityScore) }
func (r *questionResolver) PriorityClass() string         { return r.question.PriorityClass }
func (r *questionResolver) MatchedUrgencyWords() []string { return r.question.MatchedUrgencyWords }
func (r *questionResolver) PriorityDictionaryVersion() string {
	return r.question.PriorityDictionaryVersion
}
func (r *questionResolver) LastCategorizedAt() string {
	return formatTime(r.question.LastCategorizedAt)
}
func (r *questionResolver) Reply() *replyResolver {
	if r.question.Reply == nil {
		return nil
	}
	return &replyResolver{reply: *r.question.Reply}
}

type replyResolver struct{ reply domain.Reply }

func (r *replyResolver) ID() string         { return r.reply.ID }
func (r *replyResolver) QuestionID() string { return r.reply.QuestionID }
func (r *replyResolver) SellerID() string   { return r.reply.SellerID }
func (r *replyResolver) Body() string       { return r.reply.Body }
func (r *replyResolver) CreatedAt() string  { return formatTime(r.reply.CreatedAt) }

type sellerResolver struct{ seller domain.Seller }

func (r *sellerResolver) ID() string    { return r.seller.ID }
func (r *sellerResolver) Name() string  { return r.seller.Name }
func (r *sellerResolver) Email() string { return r.seller.Email }

type buyerResolver struct{ buyer domain.Buyer }

func (r *buyerResolver) ID() string    { return r.buyer.ID }
func (r *buyerResolver) Name() string  { return r.buyer.Name }
func (r *buyerResolver) Email() string { return r.buyer.Email }

type priorityQueueItemResolver struct {
	item application.QueueItem
}

func (r *priorityQueueItemResolver) Question() *questionResolver {
	return &questionResolver{question: r.item.Question}
}
func (r *priorityQueueItemResolver) Order() *orderSummaryResolver {
	return &orderSummaryResolver{order: r.item.Order}
}
func (r *priorityQueueItemResolver) Seller() *sellerResolver {
	return &sellerResolver{seller: r.item.Seller}
}
func (r *priorityQueueItemResolver) Buyer() *buyerResolver {
	return &buyerResolver{buyer: r.item.Buyer}
}
func (r *priorityQueueItemResolver) WaitingDays() int32 { return int32(r.item.WaitingDays) }

type orderSummaryResolver struct{ order domain.Order }

func (r *orderSummaryResolver) ID() string        { return r.order.ID }
func (r *orderSummaryResolver) Status() string    { return r.order.Status }
func (r *orderSummaryResolver) OrderedAt() string { return formatTime(r.order.OrderedAt) }
func (r *orderSummaryResolver) Total() float64    { return r.order.Total() }

func nullable(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func nullableTime(value *time.Time) *string {
	if value == nil {
		return nil
	}
	formatted := formatTime(*value)
	return &formatted
}

func formatTime(value time.Time) string { return value.UTC().Format(time.RFC3339Nano) }
