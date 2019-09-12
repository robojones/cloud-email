package service

import (
	"context"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/robojones/cloud-email/lib/email"
	"gotest.tools/assert"
	"testing"
)

func TestService_AddEmail(t *testing.T) {
	const (
		testEmail        = "test@email.com"
		userId    uint64 = 2342342
	)

	mockDB, mock, _ := sqlmock.New()
	mock.ExpectExec(`.*INSERT INTO emails \(id, email, user_id, created_at\).*`).
		WithArgs(sqlmock.AnyArg(), testEmail, userId).
		WillReturnResult(driver.ResultNoRows)

	s := NewService(mockDB)

	res, err := s.AddEmail(context.Background(), &email.AddEmailRequest{
		Email: testEmail,
		User: userId,
	})

	assert.NilError(t, err)
	assert.NilError(t, mock.ExpectationsWereMet())
	assert.Assert(t, 0 != res.EmailId)
	assert.Equal(t, email.AddEmailResponse_SUCCESS, res.Status)
}
