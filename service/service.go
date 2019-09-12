package service

import (
	"context"
	"github.com/robojones/cloud-email/lib/email"
)

func NewService() *Service {
	return &Service{}
}

type Service struct {}

func (s Service) AddEmail(context.Context, *email.AddEmailRequest) (*email.AddEmailResponse, error) {
	panic("implement me")
}

func (s Service) RemoveEmail(context.Context, *email.RemoveEmailRequest) (*email.RemoveEmailResponse, error) {
	panic("implement me")
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

