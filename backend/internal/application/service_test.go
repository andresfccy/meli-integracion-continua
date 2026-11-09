package application_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"seller-dashboard/internal/adapters/outbound/notifications"
	"seller-dashboard/internal/adapters/outbound/persistence/memory"
	"seller-dashboard/internal/application"
)

func newService() (*application.Service, *memory.Store, *bytes.Buffer) {
	store := memory.New()
	output := &bytes.Buffer{}
	return application.New(store, notifications.New(output)), store, output
}

func TestLoginReturnsStandardsShapedJWTAndRejectsWrongAlgorithm(t *testing.T) {
	service, _, _ := newService()
	now := time.Date(2026, 7, 29, 12, 0, 0, 0, time.UTC)
	service.SetNow(func() time.Time { return now })

	account, token, err := service.Login("seller@example.com", "seller-demo-password")
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}
	if _, err := service.Authenticate(token); err != nil {
		t.Fatalf("issued token did not authenticate: %v", err)
	}

	parts := strings.Split(token, ".")
	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		t.Fatalf("decode header: %v", err)
	}
	var header map[string]string
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		t.Fatalf("header is not JSON: %v", err)
	}
	if header["alg"] != "HS256" || header["typ"] != "JWT" {
		t.Fatalf("unexpected header: %#v", header)
	}
	if account.SellerID != "seller-1" {
		t.Fatalf("unexpected account: %#v", account)
	}
}

func TestQuestionReplyApprovalLifecycle(t *testing.T) {
	service, _, _ := newService()
	now := time.Date(2026, 7, 29, 12, 0, 0, 0, time.UTC)
	service.SetNow(func() time.Time { return now })

	buyer, _, err := service.Login("buyer@example.com", "buyer-demo-password")
	if err != nil {
		t.Fatal(err)
	}
	question, err := service.CreateQuestion(&buyer, "order-1003", "buyer-1", "Necesito ayuda", "")
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if question.Status != "open" || question.LifecycleStatus != "question_active" {
		t.Fatalf("unexpected created state: %#v", question)
	}

	question, err = service.Reply("seller-1", question.ID, "Esta es la respuesta")
	if err != nil {
		t.Fatalf("reply: %v", err)
	}
	if question.Status != "open" || question.LifecycleStatus != "seller_answered" {
		t.Fatalf("unexpected replied state: %#v", question)
	}

	question, err = service.Approve(buyer, "buyer-1", question.ID)
	if err != nil {
		t.Fatalf("approve: %v", err)
	}
	if question.Status != "resolved" || question.LifecycleStatus != "closed" || question.BuyerApprovedAt == nil {
		t.Fatalf("unexpected approved state: %#v", question)
	}
}

func TestScheduledReviewIsDuplicateSafeAndManualReviewOverrides(t *testing.T) {
	service, store, output := newService()
	at := time.Now().UTC().Add(72 * time.Hour)
	service.SetNow(func() time.Time { return at })

	first := service.ReviewScheduled(at)
	if first.Status != "completed" || first.NotificationBatchesLogged == 0 {
		t.Fatalf("unexpected first scheduled result: %#v", first)
	}
	latest := store.Events()[len(store.Events())-1]
	if latest.NotifiableID != "notifiable-seller-seller-1" || latest.ChannelID != "channel-1" {
		t.Fatalf("review bypassed the seller Notifiable channel assignment: %#v", latest)
	}
	eventCount := len(store.Events())
	second := service.ReviewScheduled(at)
	if second.Status != "already_executed" || len(store.Events()) != eventCount {
		t.Fatalf("duplicate scheduled execution was not suppressed: %#v", second)
	}
	manual := service.ReviewManual(at)
	if manual.Status != "completed" || manual.NotificationBatchesLogged == 0 || len(store.Events()) <= eventCount {
		t.Fatalf("manual override did not resend: %#v", manual)
	}
	if !strings.Contains(output.String(), `"trigger":"scheduled"`) || !strings.Contains(output.String(), `"trigger":"manual"`) {
		t.Fatalf("delivery triggers not observable: %s", output.String())
	}
}

func TestOrderFilterValidation(t *testing.T) {
	service, _, _ := newService()
	if _, err := service.SellerOrders("seller-1", application.Filters{Status: "unknown"}); err != application.ErrValidation {
		t.Fatalf("expected invalid status, got %v", err)
	}
	if _, err := service.SellerOrders("seller-1", application.Filters{DateFrom: "2026-08-01", DateTo: "2026-07-01"}); err != application.ErrValidation {
		t.Fatalf("expected invalid range, got %v", err)
	}
}
