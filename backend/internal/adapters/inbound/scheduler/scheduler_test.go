package scheduler

import (
	"context"
	"testing"
	"time"

	"seller-dashboard/internal/application"
)

type fakeReviewer struct {
	called chan time.Time
}

func (f *fakeReviewer) ReviewScheduled(at time.Time) application.ReviewResult {
	f.called <- at
	return application.ReviewResult{ReviewedAt: at, Status: "completed"}
}

func TestNextWindowIsNextMidnightUTC(t *testing.T) {
	now := time.Date(2026, 7, 29, 23, 59, 0, 0, time.FixedZone("local", -5*60*60))
	got := NextWindow(now)
	want := time.Date(2026, 7, 31, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("got %s want %s", got, want)
	}
}

func TestRunExecutesAndStopsOnCancellation(t *testing.T) {
	reviewer := &fakeReviewer{called: make(chan time.Time, 1)}
	s := New(reviewer)
	now := time.Date(2026, 7, 29, 12, 0, 0, 0, time.UTC)
	fired := make(chan time.Time, 1)
	s.SetClock(func() time.Time { return now }, func(time.Duration) <-chan time.Time { return fired })

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		s.Run(ctx)
		close(done)
	}()
	fired <- NextWindow(now)
	select {
	case <-reviewer.called:
	case <-time.After(time.Second):
		t.Fatal("scheduler did not execute")
	}
	cancel()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("scheduler did not stop after cancellation")
	}
}
