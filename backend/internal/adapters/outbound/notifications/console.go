package notifications

import (
	"encoding/json"
	"io"

	"seller-dashboard/internal/domain"
)

type Console struct{ out io.Writer }

func NewConsole(out io.Writer) *Console { return &Console{out: out} }
func (c *Console) Deliver(e domain.NotificationEvent, ch domain.NotificationChannel, trigger string) error {
	payload := map[string]any{"target": e.NotifiableID, "channelType": ch.Type, "channelId": ch.ID, "questionIds": e.QuestionIDs, "priorityClasses": e.IncludedPriorityClasses, "trigger": trigger, "timestamp": e.CreatedAt}
	return json.NewEncoder(c.out).Encode(payload)
}
