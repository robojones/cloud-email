//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/robojones/cloud-email/db"
	"github.com/robojones/cloud-email/env"
	"github.com/robojones/cloud-email/server"
	"github.com/robojones/cloud-email/service"
)

func InitServer() *server.Server {
	wire.Build(
		env.NewEnv,
		db.NewCockroachDB,
		service.NewService,
		server.NewServer,
	)

	return nil
}
