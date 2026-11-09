package notifications

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"seller-dashboard/internal/domain"
)

func TestConsoleDriverHonorsChannelStateAndType(t *testing.T) {
	var output bytes.Buffer
	driver := New(&output)
	event := domain.NotificationEvent{NotifiableID: "seller-1", QuestionIDs: []string{"q1"}, CreatedAt: time.Now()}

	if err := driver.Deliver(event, domain.NotificationChannel{Enabled: false, Type: "email"}, "scheduled"); err != nil || output.Len() != 0 {
		t.Fatalf("disabled channel should be skipped, err=%v output=%q", err, output.String())
	}
	if err := driver.Deliver(event, domain.NotificationChannel{Enabled: true, Type: "sms"}, "scheduled"); err == nil || !strings.Contains(err.Error(), "unsupported") {
		t.Fatalf("unsupported channel should fail deterministically, err=%v", err)
	}
	if err := driver.Deliver(event, domain.NotificationChannel{ID: "email-1", Enabled: true, Type: "email"}, "scheduled"); err != nil {
		t.Fatalf("email delivery failed: %v", err)
	}
	if !strings.Contains(output.String(), `"trigger":"scheduled"`) {
		t.Fatalf("missing trigger payload: %s", output.String())
	}
}
