package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/lomins/wildberriesL0/internal/config"
	"github.com/lomins/wildberriesL0/internal/models"
)

type Pg struct {
	Sqlx *sqlx.DB
	Cfg  config.Config
}

const createTableStr = `
	CREATE TABLE IF NOT EXISTS orders (
		id 		   SERIAL PRIMARY KEY,
		order_uid  VARCHAR(100) UNIQUE,
		order_data JSONB);
	CREATE INDEX IF NOT EXISTS oid ON orders(order_uid);`

func New(cfg config.Config) *Pg {
	conn, err := sqlx.Open("postgres", cfg.PgConn)
	if err != nil {
		log.Fatal("Couldn't open database: ", cfg.PgConn, err)
	}

	pg := Pg{conn, cfg}

	if _, err := pg.Sqlx.Exec(createTableStr); err != nil {
		log.Fatal("Can't create db: ", pg.Cfg.PgConn, conn, err)
	}
	return &pg
}

func (pg *Pg) Close() error {
	return pg.Sqlx.Close()
}

func (pg *Pg) InsertOrder(order models.Order) {
	q := "INSERT INTO orders (order_uid, order_data) VALUES(:order_uid, :order_data)"

	_, err := pg.Sqlx.NamedExec(q, order)
	if err != nil {
		if isDuplicateKeyError(err) {
			log.Println("pg.InsertOrder orderID already exists: ", order.ID)
			return
		}
		log.Println("pg.InsertOrder unexpected error: ", order.ID, err)
	}
}

func isDuplicateKeyError(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == "23505"
}

func (pg *Pg) GetAllOrders() (orders []models.Order) {
	query := "SELECT order_uid, order_data FROM orders"
	err := pg.Sqlx.Select(&orders, query)
	if err != nil {
		log.Println("pg.GetAllOrders failed: ", err)
	}
	return orders
}
