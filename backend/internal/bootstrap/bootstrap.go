package bootstrap

import (
	"net/http"
	"os"
	"seller-dashboard/internal/adapters/inbound/http"
	"seller-dashboard/internal/adapters/inbound/scheduler"
	"seller-dashboard/internal/adapters/outbound/notifications"
	"seller-dashboard/internal/adapters/outbound/persistence/memory"
	"seller-dashboard/internal/application"
)

type App struct {
	Handler   http.Handler
	Scheduler *scheduler.Scheduler
}

func New() App {
	store := memory.New()
	delivery := notifications.New(os.Stdout)
	service := application.New(store, delivery)
	return App{Handler: httpadapter.New(service), Scheduler: scheduler.New(service)}
}

func Handler() http.Handler {
	return New().Handler
}
