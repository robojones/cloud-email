package service

import (
	"context"
	"github.com/robojones/cloud-email/lib/email"
)

func (s Service) RemoveEmail(ctx context.Context, req *email.RemoveEmailRequest) (*email.RemoveEmailResponse, error) {
	_, err := s.db.Exec(`
		DELETE FROM emails WHERE id = $1
	`, req.EmailId)

	if err != nil {
		return nil, err
	}

	return &email.RemoveEmailResponse{
		Status: email.RemoveEmailResponse_SUCCESS,
	}, nil
}
