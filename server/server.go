package server

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/robojones/cloud-email/lib/email"
	"github.com/robojones/cloud-email/db"
	"github.com/robojones/cloud-email/env"
	"github.com/robojones/cloud-email/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewServer(e *env.Env, db db.CockroachDB, s *service.Service) *Server {
	return &Server{
		env:     e,
		db:      db,
		service: s,
	}
}

type Server struct {
	env     *env.Env
	db      db.CockroachDB
	service *service.Service
}

func (s *Server) ListenAndServe() {
	defer s.clean()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.env.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	email.RegisterEmailServer(grpcServer, s.service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(errors.Wrap(err, "error during gRPC serve"))
	}
}

func (s *Server) clean() {
	if err := s.db.Close(); err != nil {
		log.Fatal(errors.Wrap(err, "error disconnecting from database"))
	}
}
