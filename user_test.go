package user

import (
	"github.com/google/uuid"
	"testing"
)

func getUser() User {
	var pUser User = NewUser("john.doe")
	pUser.SetEmail("john.doe@tester.com")
	return pUser
}

func Test_should_get_username_from_user(t *testing.T) {
	pUser := getUser()

	expected := "john.doe"
	username := pUser.GetUsername()
	if username == nil {
		t.Errorf("Username is nil")
	}
	if expected != *username {
		t.Errorf("The Username returned is %s, while expected 'john.doe'.", *username)
	}

	expected = "john.doe.1"
	if expected == *username {
		t.Errorf("The test should fail, because the expected %s is not equals to the username", expected)
	}
}

func Test_should_get_email_from_user(t *testing.T) {
	pUser := getUser()
	var expected string = "john.doe@tester.com"
	email := pUser.GetEmail()

	if email == nil {
		t.Errorf("The email is nil")
	}
	if expected != *email {
		t.Errorf("The email address returned is %s, while expected 'john.doe@tester.com'", *email)
	}
}

func Test_should_set_userid(t *testing.T) {
	pUser := getUser()
	initId := pUser.InitId()

	if *pUser.GetUserId() != initId {
		t.Errorf("The GetUserId '%s' is not equals to the id value set '%s'.", pUser.GetUserId(), initId)
	}

	newId := IdType(uuid.New())
	pUser.SetUserId(newId)

	changedId := pUser.GetUserId()
	if changedId == nil {
		t.Errorf("Could not create a Id")
	}
	if *changedId != newId {
		t.Errorf("Could not set the id on the user")
	}
}
