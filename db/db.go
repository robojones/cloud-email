package db

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/robojones/cloud-email/env"
	"log"

	_ "github.com/lib/pq"
)

type CockroachDB = *sql.DB

// NewCockroachDB connects to the specified database and returns the client.
func NewCockroachDB(env *env.Env) CockroachDB {
	db, err := sql.Open("postgres", env.CockroachConnection)

	if err != nil {
		log.Fatal(errors.Wrap(err, "error connecting to the database"))
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error pinging the database"))
	}

	return CockroachDB(db)
}
