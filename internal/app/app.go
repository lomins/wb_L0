package app

import (
	"encoding/json"
	"log"

	"github.com/lomins/wildberriesL0/internal/cache"
	"github.com/lomins/wildberriesL0/internal/models"
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

func (app *App) ProcessNatsMessage(data []byte) {
	log.Println("proccessing Nats message...")
	o := models.Order{}

	err := json.Unmarshal(data, &o)
	if err != nil {
		log.Println("ProcessNatsMessage: error at unmarshalling the data", err)
		return
	}

	if len(o.ID) < 1 {
		log.Println("ProcessNatsMessage error: order_uid tag was not found in the input data")
		return
	}

	o.Data = data

	app.ch.Add(o.ID, o.Data)

	app.pg.InsertOrder(o)
}
