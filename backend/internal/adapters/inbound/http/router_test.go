package httpadapter_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"seller-dashboard/internal/bootstrap"
)

func request(t *testing.T, handler http.Handler, method, path, token, body string) (int, map[string]any) {
	t.Helper()
	req, err := http.NewRequest(method, path, bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)
	raw := recorder.Body.Bytes()
	var payload map[string]any
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &payload); err != nil {
			t.Fatalf("decode %s: %v", raw, err)
		}
	}
	return recorder.Code, payload
}

func login(t *testing.T, handler http.Handler, email, password string) string {
	t.Helper()
	status, payload := request(t, handler, http.MethodPost, "/api/auth/login", "", `{"email":"`+email+`","password":"`+password+`"}`)
	if status != http.StatusOK {
		t.Fatalf("login status %d: %#v", status, payload)
	}
	return payload["accessToken"].(string)
}

func TestProtectedSellerAndOperationsRoutes(t *testing.T) {
	handler := bootstrap.Handler()

	if status, _ := request(t, handler, http.MethodGet, "/api/sellers/seller-1/orders", "", ""); status != http.StatusUnauthorized {
		t.Fatalf("seller route without token returned %d", status)
	}
	buyerToken := login(t, handler, "buyer@example.com", "buyer-demo-password")
	if status, _ := request(t, handler, http.MethodGet, "/api/sellers/seller-1/orders", buyerToken, ""); status != http.StatusForbidden {
		t.Fatalf("seller route with buyer token returned %d", status)
	}
	if status, _ := request(t, handler, http.MethodGet, "/api/operations/questions", buyerToken, ""); status != http.StatusForbidden {
		t.Fatalf("operations route with buyer token returned %d", status)
	}
	opsToken := login(t, handler, "ops@example.com", "ops-demo-password")
	if status, _ := request(t, handler, http.MethodGet, "/api/operations/questions", opsToken, ""); status != http.StatusOK {
		t.Fatalf("operations route with ops token returned %d", status)
	}
}

func TestAuthorizationRequiresBearerScheme(t *testing.T) {
	handler := bootstrap.Handler()
	token := login(t, handler, "seller@example.com", "seller-demo-password")
	for _, header := range []string{token, "Basic " + token, "bearer " + token, "Bearer "} {
		req := httptest.NewRequest(http.MethodGet, "/api/sellers/seller-1/orders", nil)
		req.Header.Set("Authorization", header)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusUnauthorized {
			t.Fatalf("header %q returned %d", header, recorder.Code)
		}
	}
}

func TestGraphQLRejectsMutationsAndWrongRole(t *testing.T) {
	handler := bootstrap.Handler()
	buyerToken := login(t, handler, "buyer@example.com", "buyer-demo-password")

	status, payload := request(t, handler, http.MethodPost, "/graphql", buyerToken, `{"query":"mutation { anything }"}`)
	if status != http.StatusOK || payload["errors"] == nil {
		t.Fatalf("mutation was not rejected: status=%d payload=%#v", status, payload)
	}
	status, payload = request(t, handler, http.MethodPost, "/graphql", buyerToken, `{"query":"query { operationsPriorityQueue { questions { waitingDays } } }"}`)
	if status != http.StatusOK || payload["errors"] == nil {
		t.Fatalf("operations GraphQL with buyer token was not rejected: status=%d payload=%#v", status, payload)
	}
}

func TestGraphQLExecutesSchemaVariablesAliasesAndFragments(t *testing.T) {
	handler := bootstrap.Handler()
	token := login(t, handler, "seller@example.com", "seller-demo-password")
	body := `{
		"query":"query Seller($sellerId:String!,$filters:OrderFilters){ current: viewer { ...Identity } sellerDashboard(sellerId:$sellerId,filters:$filters){ orders { id status buyer { name } } } } fragment Identity on Account { id roles }",
		"variables":{"sellerId":"seller-1","filters":{"status":"shipped"}}
	}`
	status, payload := request(t, handler, http.MethodPost, "/graphql", token, body)
	if status != http.StatusOK || payload["errors"] != nil {
		t.Fatalf("GraphQL query failed: status=%d payload=%#v", status, payload)
	}
	data := payload["data"].(map[string]any)
	if data["current"] == nil {
		t.Fatalf("alias/fragment result missing: %#v", data)
	}
	orders := data["sellerDashboard"].(map[string]any)["orders"].([]any)
	if len(orders) != 1 || orders[0].(map[string]any)["status"] != "shipped" {
		t.Fatalf("filter variables were not applied: %#v", orders)
	}
}

func TestGraphQLValidatesUnknownFields(t *testing.T) {
	handler := bootstrap.Handler()
	token := login(t, handler, "seller@example.com", "seller-demo-password")
	status, payload := request(t, handler, http.MethodPost, "/graphql", token, `{"query":"query { viewer { fieldThatDoesNotExist } }"}`)
	if status != http.StatusOK || payload["errors"] == nil {
		t.Fatalf("unknown field was not rejected: status=%d payload=%#v", status, payload)
	}
}

func TestGraphQLExposesReadOnlySchemaThroughIntrospection(t *testing.T) {
	handler := bootstrap.Handler()
	token := login(t, handler, "seller@example.com", "seller-demo-password")
	status, payload := request(t, handler, http.MethodPost, "/graphql", token, `{"query":"query { __schema { queryType { name } mutationType { name } } }"}`)
	if status != http.StatusOK || payload["errors"] != nil {
		t.Fatalf("introspection failed: status=%d payload=%#v", status, payload)
	}
	schema := payload["data"].(map[string]any)["__schema"].(map[string]any)
	if schema["queryType"].(map[string]any)["name"] != "Query" {
		t.Fatalf("query root missing: %#v", schema)
	}
	if schema["mutationType"] != nil {
		t.Fatalf("read-only schema exposed a mutation root: %#v", schema)
	}
}

func TestManualReviewValidatesTimestamp(t *testing.T) {
	handler := bootstrap.Handler()
	token := login(t, handler, "ops@example.com", "ops-demo-password")
	status, _ := request(t, handler, http.MethodPost, "/api/operations/priority-review", token, `{"reviewedAt":"not-a-date"}`)
	if status != http.StatusBadRequest {
		t.Fatalf("invalid reviewedAt returned %d", status)
	}
}

func TestBuyerSellerApprovalFlow(t *testing.T) {
	handler := bootstrap.Handler()
	buyerToken := login(t, handler, "buyer@example.com", "buyer-demo-password")
	status, created := request(t, handler, http.MethodPost, "/api/buyers/buyer-1/orders/order-1003/questions", buyerToken, `{"body":"Necesito ayuda con el monitor"}`)
	if status != http.StatusCreated {
		t.Fatalf("create returned %d: %#v", status, created)
	}
	question := created["question"].(map[string]any)
	questionID := question["id"].(string)

	sellerToken := login(t, handler, "seller@example.com", "seller-demo-password")
	status, replied := request(t, handler, http.MethodPost, "/api/sellers/seller-1/questions/"+questionID+"/reply", sellerToken, `{"body":"Revisamos tu caso"}`)
	if status != http.StatusCreated {
		t.Fatalf("reply returned %d: %#v", status, replied)
	}
	if replied["question"].(map[string]any)["lifecycleStatus"] != "seller_answered" {
		t.Fatalf("unexpected reply state: %#v", replied)
	}

	status, approved := request(t, handler, http.MethodPost, "/api/buyers/buyer-1/questions/"+questionID+"/approve-answer", buyerToken, `{}`)
	if status != http.StatusOK {
		t.Fatalf("approve returned %d: %#v", status, approved)
	}
	result := approved["question"].(map[string]any)
	if result["status"] != "resolved" || result["lifecycleStatus"] != "closed" {
		t.Fatalf("unexpected approval state: %#v", result)
	}
}

func TestBackendQuestionCreationContract(t *testing.T) {
	handler := bootstrap.Handler()
	status, created := request(t, handler, http.MethodPost, "/api/orders/order-1002/questions", "", `{"buyerId":"buyer-2","body":"Necesito otra aclaración","orderItemId":"item-1003"}`)
	if status != http.StatusCreated {
		t.Fatalf("backend question creation returned %d: %#v", status, created)
	}
	question := created["question"].(map[string]any)
	if question["buyerId"] != "buyer-2" || question["orderItemId"] != "item-1003" {
		t.Fatalf("unexpected backend-created question: %#v", question)
	}
}

func TestSellerFiltersAndCORS(t *testing.T) {
	handler := bootstrap.Handler()
	token := login(t, handler, "seller@example.com", "seller-demo-password")
	status, payload := request(t, handler, http.MethodGet, "/api/sellers/seller-1/orders?status=unknown", token, "")
	if status != http.StatusBadRequest || payload["error"] == nil {
		t.Fatalf("invalid filter returned %d: %#v", status, payload)
	}

	req := httptest.NewRequest(http.MethodOptions, "/api/auth/login", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusNoContent || recorder.Header().Get("Access-Control-Allow-Origin") != "http://localhost:3000" {
		t.Fatalf("unexpected preflight: status=%d headers=%v", recorder.Code, recorder.Header())
	}
}
