package service

import (
	"context"
	"github.com/omeid/pgerror"
	"github.com/robojones/cloud-email/db"
	"github.com/robojones/cloud-email/lib/email"
	"github.com/robojones/iid"
)

func NewService(db db.CockroachDB) *Service {
	return &Service{
		db: db,
	}
}

type Service struct {
	db db.CockroachDB
}

func (s Service) AddEmail(ctx context.Context, req *email.AddEmailRequest) (*email.AddEmailResponse, error) {
	id := iid.New()

	_, err := s.db.Exec(`
		INSERT INTO emails (id, email, user_id, created_at) VALUES ($1, $2, $3, DEFAULT)
	`, id.Int(), req.Email, req.User)

	if e := pgerror.UniqueViolation(err); e != nil {
		return &email.AddEmailResponse{
			Status:               email.AddEmailResponse_DUPLICATE_EMAIL,
		}, nil
	}

	if err != nil {
		return nil, err
	}

	return &email.AddEmailResponse{
		Status: email.AddEmailResponse_SUCCESS,
		EmailId: id.Uint64(),
	}, nil
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

