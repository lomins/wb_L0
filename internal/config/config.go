package config

type Config struct {
	WebSerberAddr string
	PgConn        string
	NatsURL       string
	NatsCluster   string
	NatsDurable   string
	NatsSubject   string
}

func New() *Config {
	return &Config{
		WebSerberAddr: "http://localhost:8080",
		PgConn:        "user=postgres password=7070 dbname=wbL0 port=5432 sslmode=disable",
		NatsURL:       "nats://localhost:4222",
		NatsCluster:   "my-cluster",
		NatsDurable:   "my-durable",
		NatsSubject:   "wb-orders",
	}
}
