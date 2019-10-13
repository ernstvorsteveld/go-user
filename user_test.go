package user

import (
	"github.com/google/uuid"
	"testing"
)

func getUser() User {
	var user User = NewUser("john.doe")
	user.SetEmail("john.doe@tester.com")
	return user
}

func Test_should_get_username_from_user(t *testing.T) {
	pUser := getUser()

	expected := "john.doe"
	username := pUser.GetUsername()
	if username == "" {
		t.Errorf("Username is empty")
	}
	if expected != username {
		t.Errorf("The Username returned is %s, while expected 'john.doe'.", username)
	}

	expected = "john.doe.1"
	if expected == username {
		t.Errorf("The test should fail, because the expected %s is not equals to the username", expected)
	}
}

func Test_should_get_email_from_user(t *testing.T) {
	user := getUser()
	var expected string = "john.doe@tester.com"
	email := user.GetEmail()

	if email == "" {
		t.Errorf("The email is empty")
	}
	if expected != email {
		t.Errorf("The email address returned is %s, while expected 'john.doe@tester.com'", email)
	}
}

func Test_should_set_userid(t *testing.T) {
	user := getUser()
	initId := user.InitId()

	if user.GetUserId() != initId {
		t.Errorf("The GetUserId '%s' is not equals to the id value set '%s'.", user.GetUserId(), initId)
	}

	newId := IdType(uuid.New())
	user.SetUserId(newId)

	changedId := user.GetUserId()
	if changedId != newId {
		t.Errorf("Could not set the id on the user")
	}
}

func Test_methods_on_user(t *testing.T) {
	user := getUser()

	expectedEmail := "test.email@host.com"
	user.SetEmail(expectedEmail)
	foundEmail := user.GetEmail()

	if foundEmail != expectedEmail {
		t.Errorf("The expected email %s is not equals to the found email %s", expectedEmail, foundEmail)
	}

	expectedUsername := "test.username"
	user.SetUsername(expectedUsername)
	foundUsername := user.GetUsername()

	if foundUsername != expectedUsername {
		t.Errorf("The expected username %s is not equals to the found username %s.", expectedUsername, foundUsername)
	}

}
