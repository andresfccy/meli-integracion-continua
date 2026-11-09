package domain

import (
	"testing"
	"time"
)

func TestEvaluatePriorityCriticalPhraseWins(t *testing.T) {
	now := time.Date(2026, 7, 18, 0, 0, 0, 0, time.UTC)
	score, class, words := EvaluatePriority("Urgent refund because it never arrived", now.Add(-72*time.Hour), now, 200)
	if class != "Critical" || score < 80 || len(words) == 0 {
		t.Fatalf("expected critical classification, got %s (%d), words=%v", class, score, words)
	}
}

func TestOrderTotalUsesPurchasePrice(t *testing.T) {
	o := Order{Items: []OrderItem{{Quantity: 2, UnitPrice: 19.5}, {Quantity: 1, UnitPrice: 10}}}
	if got := o.Total(); got != 49 {
		t.Fatalf("expected 49, got %v", got)
	}
}

func TestEvaluatePriorityThresholdsAndCaps(t *testing.T) {
	now := time.Date(2026, 7, 29, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name  string
		body  string
		age   time.Duration
		value float64
		class string
	}{
		{"low", "hello", 0, 0, "Low"},
		{"medium", "hello", 6 * 24 * time.Hour, 0, "Medium"},
		{"high", "urgent", 0, 0, "High Priority"},
		{"critical", "never arrived", 0, 0, "Critical"},
		{"future clamps waiting", "hello", -24 * time.Hour, 0, "Low"},
		{"score caps", "never arrived", 20 * 24 * time.Hour, 10000, "Critical"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score, class, _ := EvaluatePriority(tt.body, now.Add(-tt.age), now, tt.value)
			if class != tt.class {
				t.Fatalf("class=%s score=%d", class, score)
			}
			if score > 100 {
				t.Fatalf("score exceeded cap: %d", score)
			}
		})
	}
}

func TestNotifiableFactoryBuildsSellerAndOperationsTargets(t *testing.T) {
	factory := NewNotifiableFactory()
	sellerChannels := []string{"seller-email"}
	seller, err := factory.FromSeller(Seller{ID: "seller-1", NotificationChannelIDs: sellerChannels})
	if err != nil {
		t.Fatal(err)
	}
	member, err := factory.FromOpsTeamMember(OpsTeamMember{ID: "ops-1", NotificationChannelIDs: []string{"ops-email"}})
	if err != nil {
		t.Fatal(err)
	}
	if seller.ID != "notifiable-seller-seller-1" || seller.ObjectType != "seller" || seller.NotificationChannelIDs[0] != "seller-email" {
		t.Fatalf("unexpected seller target: %#v", seller)
	}
	if member.ID != "notifiable-opsTeamMember-ops-1" || member.ObjectType != "opsTeamMember" {
		t.Fatalf("unexpected Operations target: %#v", member)
	}

	sellerChannels[0] = "changed-outside"
	if seller.NotificationChannelIDs[0] != "seller-email" {
		t.Fatal("factory retained the caller's mutable channel slice")
	}
	if _, err := factory.FromSeller(Seller{}); err == nil {
		t.Fatal("seller without id should be rejected")
	}
	if _, err := factory.FromOpsTeamMember(OpsTeamMember{}); err == nil {
		t.Fatal("Operations member without id should be rejected")
	}
}
