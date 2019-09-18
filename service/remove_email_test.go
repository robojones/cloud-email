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

func TestService_RemoveEmail(t *testing.T) {

	mockDB, mock, _ := sqlmock.New()
	mock.ExpectExec(`DELETE FROM emails WHERE id="2342342" `).WithArgs(sqlmock.AnyArg(), testEmail, userId).
		WillReturnResult(driver.ResultNoRows)

	s := NewService(mockDB)

	res, err := s.RemoveEmail(context.Background(), &email.RemoveEmailRequest{
		User:                 0,
		EmailId:              0,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})

	assert.NilError(t, err)
	assert.NilError(t, mock.ExpectationsWereMet())
	assert.Equal(t, email.AddEmailResponse_SUCCESS, res.Status)
}