package env

import (
	"github.com/caarlos0/env/v6"
	"log"
)

func NewEnv() *Env {
	e := &Env{}
	if err := env.Parse(e); err != nil {
		log.Fatal(err)
	}
	return e
}

type Env struct {
	// CockroachConnection is the connection string used to connect to the CockroachDB cluster.
	// Example value: "host=localhost port=26257 user=email_test dbname=email_test sslmode=disable"
	CockroachConnection string `env:"COCKROACH_CONNECTION,required"`

	// Port that the gRPC server listens to.
	Port int `env:"PORT,required"`
}
