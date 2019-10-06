package user

import (
	"github.com/google/uuid"
	"testing"
)

func getUser() *UserIdentity {
	user := UserIdentity{
		User: UserType{
			Id:       uuid.New(),
			Username: "john.doe",
			Email:    "john.doe@tester.com",
			Name: NameType{
				First: "John",
				Last:  "Doe",
			},
		},
	}
	return &user
}

func Test_should_get_username_from_user(t *testing.T) {
	pUser := getUser()

	var expected UsernameType = "john.doe"
	username := GetUsername(pUser)
	if expected != username {
		t.Errorf("The Username returned is %s, while expected 'john.doe'.", username)
	}

	expected = "john.doe.1"
	if expected == username {
		t.Errorf("The test should fail, because the expected %s is not equals to the username", expected)
	}
}

func Test_should_get_email_from_user(t *testing.T) {
	pUser := getUser()
	var expected EmailType = "john.doe@tester.com"
	email := GetEmail(pUser)
	if expected != email {
		t.Errorf("The email address returned is %s, while expected 'john.doe@tester.com'", email)
	}
}
