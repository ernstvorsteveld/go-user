package user

import (
	"github.com/google/uuid"
	"testing"
)

func Test_should_get_username_from_user(t *testing.T) {
	pUser := &UserIdentity{
		User: UserType{
			Id:       uuid.New(),
			Username: "john.doe",
			Name: NameType{
				First: "John",
				Last:  "Doe",
			},
		},
	}

	var expected UsernameType = "john.doe"
	username := getUsername(pUser)
	if expected != username {
		t.Errorf("The Username returned is %s, while expected 'john.doe'.", username)
	}

	expected = "john.doe.1"
	if expected == username {
		t.Errorf("The test should fail, because the expected %s is not equals to the username", expected)
	}
}
