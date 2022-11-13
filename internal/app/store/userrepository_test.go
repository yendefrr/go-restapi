package store_test

import (
	"github.com/stretchr/testify/assert"
	"go/rest-api/internal/app/model"
	"go/rest-api/internal/app/store"

	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser())
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@yendefr.xyz"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	_, err = s.User().Create(model.TestUser())
	if err != nil {
		t.Fatal(err)
	}
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
