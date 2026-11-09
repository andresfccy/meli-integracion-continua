package scheduler

import (
	"context"
	"log"
	"time"

	"seller-dashboard/internal/application"
)

type Reviewer interface {
	ReviewScheduled(time.Time) application.ReviewResult
}

type Scheduler struct {
	reviewer Reviewer
	now      func() time.Time
	after    func(time.Duration) <-chan time.Time
	logf     func(string, ...any)
}

func New(reviewer Reviewer) *Scheduler {
	return &Scheduler{reviewer: reviewer, now: time.Now, after: time.After, logf: log.Printf}
}

func (s *Scheduler) SetClock(now func() time.Time, after func(time.Duration) <-chan time.Time) {
	s.now = now
	s.after = after
}

func NextWindow(now time.Time) time.Time {
	now = now.UTC()
	return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC)
}

func (s *Scheduler) Run(ctx context.Context) {
	for {
		now := s.now().UTC()
		next := NextWindow(now)
		select {
		case <-ctx.Done():
			return
		case <-s.after(next.Sub(now)):
			s.execute(next)
		}
	}
}

func (s *Scheduler) execute(at time.Time) {
	defer func() {
		if recovered := recover(); recovered != nil {
			s.logf("scheduled priority review recovered panic: %v", recovered)
		}
	}()
	result := s.reviewer.ReviewScheduled(at)
	s.logf("scheduled priority review window=%s status=%s recategorized=%d batches=%d",
		result.ReviewWindow, result.Status, result.Recategorized, result.NotificationBatchesLogged)
}
