package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	var id int64 = 1
	var name string = "Nakata"
	var email string = "you@example.com"
	user := NewUser(id, name, email)

	if user == nil {
		t.Errorf("failed NewUser()")
	}

	assert.Equal(t, id, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)

	t.Logf("user: %p", user)
	t.Logf("user.ID: %d", user.ID)
	t.Logf("user.Name: %s", user.Name)
	t.Logf("user.Email: %s", user.Email)
}
