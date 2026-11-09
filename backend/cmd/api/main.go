package main

import (
	"context"
	"log"
	"net/http"
	"seller-dashboard/internal/bootstrap"
)

func main() {
	app := bootstrap.New()
	go app.Scheduler.Run(context.Background())
	log.Println("seller dashboard API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", app.Handler))
}
