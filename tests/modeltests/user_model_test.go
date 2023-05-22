package modeltests

import (
	"log"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllUsers(t *testing.T) {
	err := resfreshUserTables()
	if err != nil {
		log.Fatal(err)
	}

	err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v", err)
	}

	assert.Equal(t, len(*users), 2)
}
