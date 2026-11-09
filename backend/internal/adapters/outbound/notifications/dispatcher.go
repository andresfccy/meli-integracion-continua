package notifications

import (
	"fmt"
	"io"

	"seller-dashboard/internal/domain"
)

type Driver interface {
	Deliver(domain.NotificationEvent, domain.NotificationChannel, string) error
}

type Dispatcher struct {
	drivers map[string]Driver
}

// New builds the MVP notification delivery registry. The email channel intent
// is currently simulated by the console driver; adding a real email, Slack, or
// SMS driver only requires registering another channel implementation here.
func New(out io.Writer) *Dispatcher {
	dispatcher := &Dispatcher{drivers: map[string]Driver{}}
	dispatcher.Register("email", NewConsole(out))
	return dispatcher
}

func (d *Dispatcher) Register(channelType string, driver Driver) {
	d.drivers[channelType] = driver
}

func (d *Dispatcher) Deliver(event domain.NotificationEvent, channel domain.NotificationChannel, trigger string) error {
	if !channel.Enabled {
		return nil
	}
	driver, ok := d.drivers[channel.Type]
	if !ok {
		return fmt.Errorf("unsupported notification channel: %s", channel.Type)
	}
	return driver.Deliver(event, channel, trigger)
}
