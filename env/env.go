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
	// CockroachURL is the connection string used to connect to the CockroachDB cluster.
	// Example value: "postgresql://email_test@localhost:26257/email_test"
	CockroachURL string `env:"COCKROACH_URL,required"`

	// Port that the gRPC server listens to.
	Port int `env:"PORT,required"`

	// Client side hashing salt seed.
	SaltSeed string `env:"SALT_SEED,required"`
}
