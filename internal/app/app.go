package app

import (
	"github.com/lomins/wildberriesL0/internal/cache"
	"github.com/lomins/wildberriesL0/internal/postgres"
)

type App struct {
	pg *postgres.Pg
	ch *cache.Cache
}

func New(pg *postgres.Pg, ch *cache.Cache) *App {
	return &App{pg, ch}
}

func (a *App) RestoreCacheDataFromPg() {
	orders := a.pg.GetAllOrders()
	a.ch.AddSet(orders)
}

func (a *App) ProccessNatsMessage() {

}
