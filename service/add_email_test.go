package service

import (
	"context"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/robojones/cloud-email/db"
	"github.com/robojones/cloud-email/env"
	"github.com/robojones/cloud-email/lib/email"
	"github.com/robojones/cloud-email/test_util"
	"gotest.tools/assert"
	"testing"
)

func TestService_GetEmail_Error(t *testing.T) {
	const (
		testEmail        = "test@email.com"
		userId    uint64 = 2342342
	)

	mockDB, mock, _ := sqlmock.New()
	mock.ExpectExec(`SELECT FROM emails*`).WithArgs(sqlmock.AnyArg(), testEmail, userId).
		WillReturnResult(driver.ResultNoRows)

	s := NewService(mockDB)

	res, err := s.AddEmail(context.Background(), &email.AddEmailRequest{
		Email: testEmail,
		User:  userId,
	})

	assert.NilError(t, err)
	assert.NilError(t, mock.ExpectationsWereMet())
	assert.Assert(t, 0 != res.EmailId)
	assert.Equal(t, email.AddEmailResponse_SUCCESS, res.Status)
}

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
		User:  userId,
	})

	assert.NilError(t, err)
	assert.NilError(t, mock.ExpectationsWereMet())
	assert.Assert(t, 0 != res.EmailId)
	assert.Equal(t, email.AddEmailResponse_SUCCESS, res.Status)
}

func TestService_AddEmail_Error (t *testing.T) {
	const (
		testEmail        = "test@email.com"
		userId    uint64 = 2342342
	)

	mockDB, mock, _ := sqlmock.New();

	s := NewService(mockDB)

	var ConnectionError = &pq.Error{
		Code: "08006",
	}

	mock.ExpectExec(`INSERT INTO emails .*`).WillReturnError(ConnectionError)

	_, err := s.AddEmail(context.Background(), &email.AddEmailRequest{
		Email: testEmail,
		User:  userId,
	})
	assert.Equal(t, ConnectionError, err)
}

func TestService_AddEmail_UniqueViolationError(t *testing.T) {
	const (
		testEmail        = "test@email.com"
		userId    uint64 = 2342342
	)

	var UniqueViolationError = &pq.Error{
		Code: "23505",
	}

	mockDB, mock, _ := sqlmock.New()
	mock.ExpectExec(`.*INSERT INTO emails \(id, email, user_id, created_at\).*`).
		WithArgs(sqlmock.AnyArg(), testEmail, userId).
		WillReturnError(UniqueViolationError)

	s := NewService(mockDB)

	res, err := s.AddEmail(context.Background(), &email.AddEmailRequest{
		Email: testEmail,
		User:  userId,
	})

	assert.NilError(t, err)
	assert.Equal(t, email.AddEmailResponse_DUPLICATE_EMAIL, res.Status)
}

func TestService_AddEmail_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	const (
		testEmail        = "test@email.com"
		userId    uint64 = 2342342
	)

	pg := db.NewCockroachDB(env.NewEnv())
	c := NewService(pg)

	defer test_util.ClearTable(t, pg, "emails")

	successResult, err := c.AddEmail(context.Background(), &email.AddEmailRequest{
		User:  userId,
		Email: testEmail,
	})

	assert.NilError(t, err)
	assert.Equal(t, email.AddEmailResponse_SUCCESS, successResult.Status)
	assert.Assert(t, 0 != successResult.EmailId)

	errorResult, err := c.AddEmail(context.Background(), &email.AddEmailRequest{
		User:  userId,
		Email: testEmail,
	})

	assert.NilError(t, err)
	assert.Equal(t, email.AddEmailResponse_DUPLICATE_EMAIL, errorResult.Status)
}
