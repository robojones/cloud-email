package test_util

import (
	"github.com/robojones/cloud-email/db"
	"gotest.tools/assert"
	"testing"
)

// ClearTable removes all rows from a specified table.
func ClearTable(t *testing.T, db db.CockroachDB, table string) {
	_, err := db.Exec("DELETE FROM " + table)

	assert.NilError(t, err)
}
