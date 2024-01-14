package main

import (
	"github.com/lomins/wildberriesL0/internal/app"
	"github.com/lomins/wildberriesL0/internal/cache"
	"github.com/lomins/wildberriesL0/internal/config"
	"github.com/lomins/wildberriesL0/internal/nats"
	"github.com/lomins/wildberriesL0/internal/postgres"
	"github.com/lomins/wildberriesL0/internal/webserver"
)

func main() {
	cfg := config.New()

	ch := cache.New()

	pg := postgres.New(cfg)
	defer pg.Close()

	app := app.New(pg, ch)
	app.RestoreCacheDataFromPg()

	sub := nats.New(cfg, "subscriber")
	sub.SubscribeOnSubject(app)
	defer sub.Close()

	srv := webserver.New(cfg, ch)
	go srv.ShutdownOnSignal()
	srv.Launch()
}
