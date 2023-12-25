package nats

import (
	"log"

	"github.com/lomins/wildberriesL0/internal/config"
	"github.com/nats-io/stan.go"
)

type NatsConn struct {
	cfg  config.Config
	conn stan.Conn
}

func New(cfg config.Config, clientID string) *NatsConn {
	sc, err := stan.Connect(cfg.NatsCluster, clientID, stan.NatsURL(cfg.NatsURL))
	if err != nil {
		log.Fatal("Can't connect to NATS: ", err)
	}
	log.Println("Connected to NATS, clientID = ", clientID)
	return &NatsConn{cfg: cfg, conn: sc}
}

func (n *NatsConn) Publish(data []byte) {
	log.Println("Publishing message to NATS...")
	n.conn.Publish(n.cfg.NatsSubject, data)
}
