package service

import (
	"context"
	"github.com/robojones/cloud-email/lib/email"
	"github.com/robojones/iid"
)

func (s Service) AddEmail(ctx context.Context, req *email.AddEmailRequest) (*email.AddEmailResponse, error) {
	id := iid.New()

	_, err := s.db.Exec(`
		INSERT INTO emails (id, email, user_id, created_at) VALUES ($1, $2, $3, DEFAULT)
	`, id.Int(), req.Email, req.User)

	if err != nil {
		return nil, err
	}

	return &email.AddEmailResponse{
		Status: email.AddEmailResponse_SUCCESS,
		EmailId: id.Uint64(),
	}, nil
}
