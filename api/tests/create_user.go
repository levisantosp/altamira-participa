package tests

import (
	"testing"

	"github.com/levisantosp/altamira-participa/api/db"
	"github.com/levisantosp/altamira-participa/api/ent/generated"
)

func CreateUser(t *testing.T) *generated.User {
	return db.Client.User.
		Create().
		SetUsername("test-user").
		SetDisplayName("Test User").
		SetEmail("testuser@email.com").
		SaveX(t.Context())
}
