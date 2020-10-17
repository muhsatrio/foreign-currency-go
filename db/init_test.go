package db_test

import (
	"foreign-currency-go/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	err := db.Init()
	assert.NoError(t, err)
}
