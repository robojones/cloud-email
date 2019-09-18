package service

import (
	"context"
	"github.com/robojones/cloud-email/db"
	"github.com/robojones/cloud-email/lib/email"
)

func NewService(db db.CockroachDB) *Service {
	return &Service{
		db: db,
	}
}

type Service struct {
	db db.CockroachDB
}

func (s Service) GetEmails(context.Context, *email.GetEmailsRequest) (*email.GetEmailsResponse, error) {
	panic("implement me")
}

func (s Service) ValidateEmail(context.Context, *email.ValidateEmailRequest) (*email.ValidateEmailResponse, error) {
	panic("implement me")
}

func (s Service) SetPrimaryEmail(context.Context, *email.SetPrimaryEmailRequest) (*email.SetPrimaryEmailResponse, error) {
	panic("implement me")
}

func (s Service) GetUser(context.Context, *email.GetUserRequest) (*email.GetUserResponse, error) {
	panic("implement me")
}

