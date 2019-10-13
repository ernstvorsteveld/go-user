package repository

import (
	models "github.com/ernstvorsteveld/go-user"
	"reflect"
	"testing"
)

func Test_create_and_retrieve_user(t *testing.T) {
	userRepository := NewMapUserRepository()
	userMap := createUsers(userRepository)

	john := userMap[JohnUsername]
	userFromGet, errorFromGet := userRepository.GetById(john.GetUserId())
	if errorFromGet != nil {
		t.Errorf("Error returned by user repository while getting a user.")
	}
	if JohnUsername != userFromGet.GetUsername() {
		t.Errorf("The usernames are different, expected %s, but got %s.", JohnUsername, userFromGet.GetUsername())
	}

	same := reflect.DeepEqual(john, userFromGet)
	if !same {
		t.Errorf("The user retrieved %s differs from the created user %s.", john.GetUsername(), userFromGet.GetUsername())
	}

	_, error := userRepository.GetByLoginCode(JohnUsername)
	if error != nil {
		t.Errorf("Error occurred: %s", error)
	}
}

func Test_create_update_and_delete_user(t *testing.T) {
	userRepository := NewMapUserRepository()
	userMap := createUsers(userRepository)

	expected := "username-update"
	user := models.NewUser(expected)
	id := userMap[JohnUsername].GetUserId()

	if len(id) <= 0 {
		t.Errorf("Id is nil.")
	}
	user.SetUserId(id)
	userRepository.Update(user)

	updatedUser, _ := userRepository.GetById(id)

	if updatedUser.GetUsername() != expected {
		t.Errorf("Updating of the username did not work, got %s, expected %s.", updatedUser.GetUsername(), expected)
	}
}
