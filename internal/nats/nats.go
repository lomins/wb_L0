package nats

import (
	"log"

	"github.com/lomins/wildberriesL0/internal/config"
	"github.com/nats-io/stan.go"
)

type NatsConn struct {
	cfg  config.Config
	conn stan.Conn
	sub  stan.Subscription
	proc processor
}

type processor interface {
	ProcessNatsMessage(data []byte)
}

func New(cfg config.Config, clientID string) *NatsConn {
	sc, err := stan.Connect(cfg.NatsCluster, clientID, stan.NatsURL(cfg.NatsURL))
	if err != nil {
		log.Fatal("Can't connect to NATS: ", err)
	}
	log.Println("Connected to NATS, clientID = ", clientID)
	return &NatsConn{cfg: cfg, conn: sc}
}

func (n *NatsConn) SubscribeOnSubject(next processor) {
	n.proc = next

	ss, err := n.conn.Subscribe(
		n.cfg.NatsSubject,
		n.recieveNatsMsg,
		stan.DurableName(n.cfg.NatsDurable))
	if err != nil {
		log.Fatal("error at subscribing to nats", err, n.cfg)
	}
	n.sub = ss
}

func (n *NatsConn) Publish(data []byte) {
	log.Println("Publishing message to NATS...")
	n.conn.Publish(n.cfg.NatsSubject, data)
}

func (n *NatsConn) Close() {
	if n.sub != nil {
		log.Println("closing nats subscription", n.sub.Close())
	}
	log.Println("closing nats connection", n.conn.Close())
}

func (n *NatsConn) recieveNatsMsg(m *stan.Msg) {
	log.Println("got new msg from nats", m.Size(), m.Timestamp)
	n.proc.ProcessNatsMessage(m.Data)
}
