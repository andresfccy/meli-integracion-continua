package memory

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestSeedIncludesDocumentedEntitiesAndNotifiableTargets(t *testing.T) {
	store := New()
	if len(store.Products()) == 0 || len(store.OpsTeamMembers()) == 0 || len(store.Notifiables()) < 3 {
		t.Fatalf("incomplete seed: products=%d ops=%d notifiables=%d", len(store.Products()), len(store.OpsTeamMembers()), len(store.Notifiables()))
	}
	var open, resolved bool
	for _, question := range store.Questions() {
		open = open || question.Status == "open"
		resolved = resolved || question.Status == "resolved"
	}
	if !open || !resolved {
		t.Fatalf("seed must demonstrate open and resolved lifecycles: open=%v resolved=%v", open, resolved)
	}
}

func TestReadsReturnDeepCopies(t *testing.T) {
	store := New()
	orders := store.Orders()
	orders[0].Items[0].ProductName = "mutated"
	questions := store.Questions()
	questions[0].MatchedUrgencyWords[0] = "mutated"
	sellers := store.Sellers()
	sellers[0].NotificationChannelIDs[0] = "mutated"
	events := store.Events()
	events[0].QuestionIDs[0] = "mutated"

	if store.Orders()[0].Items[0].ProductName == "mutated" ||
		store.Questions()[0].MatchedUrgencyWords[0] == "mutated" ||
		store.Sellers()[0].NotificationChannelIDs[0] == "mutated" ||
		store.Events()[0].QuestionIDs[0] == "mutated" {
		t.Fatal("a read exposed mutable store-owned data")
	}
}

func TestAcquireReviewWindowIsThreadSafeAndSingleWinner(t *testing.T) {
	store := New()
	var winners int32
	var group sync.WaitGroup
	for range 20 {
		group.Add(1)
		go func() {
			defer group.Done()
			if store.AcquireReviewWindow("2026-07-29") {
				atomic.AddInt32(&winners, 1)
			}
		}()
	}
	group.Wait()
	if winners != 1 {
		t.Fatalf("expected one winner, got %d", winners)
	}
}
