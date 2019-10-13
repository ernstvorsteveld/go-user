package repository

import (
	models "github.com/ernstvorsteveld/go-user"
	"reflect"
	"testing"
)

func Test_create_and_retrieve_user(t *testing.T) {
	userRepository := NewMapUserRepository()
	userMap := createUsers(userRepository, t)

	pJohn := userMap["john.doe"]
	pUserFromGet, errorFromGet := userRepository.GetById(pJohn.GetUserId())
	if errorFromGet != nil {
		t.Errorf("Error returned by user repository while getting a user.")
	}
	if "john.doe" != pUserFromGet.GetUsername() {
		t.Errorf("The usernames are different, expected 'john.doe', but got %s.", pJohn.GetUsername())
	}

	same := reflect.DeepEqual(userMap["john.doe"], pUserFromGet)
	if !same {
		t.Errorf("The user retrieved %s differs from the created user %s.", pJohn.GetUsername(), pUserFromGet.GetUsername())
	}

	_, error := userRepository.GetByLoginCode("john.doe")
	if error != nil {
		t.Errorf("Error occurred: %s", error)
	}
}

func createUsers(userRepository UserRepository, t *testing.T) map[string]models.User {
	userMap := make(map[string]models.User)

	username := "john.doe"
	user := models.NewUser(username)
	uuid, error := userRepository.Create(user)
	validateUserCreate(error, uuid, t)
	userMap[username] = user
	if userMap[username].GetUsername() != username {
		t.Errorf("1.Username is incoorect: %s, expected %s.", userMap[username].GetUsername(), username)
	}

	username = "sander.kuiper"
	user = models.NewUser(username)
	uuid, error = userRepository.Create(user)
	validateUserCreate(error, uuid, t)
	userMap[username] = user
	if userMap[username].GetUsername() != username {
		t.Errorf("2.Username is incoorect: %s, expected %s.", userMap[username].GetUsername(), username)
	}

	username = "annelies.gerritsen"
	user = models.NewUser(username)
	uuid, error = userRepository.Create(user)
	validateUserCreate(error, uuid, t)
	userMap[username] = user
	if userMap[username].GetUsername() != username {
		t.Errorf("3.Username is incoorect: %s, expected %s.", userMap[username].GetUsername(), username)
	}

	return userMap
}

func validateUserCreate(error error, id models.IdType, t *testing.T) {
	if error != nil {
		t.Errorf("Error returned by user repository while creating a user.")
	}
	if len(id) <= 0 {
		t.Errorf("When Creating a user in the repository, expected a UUID value.")
	}

}

func Test_create_update_and_delete_user(t *testing.T) {
	userRepository := NewMapUserRepository()
	userMap := createUsers(userRepository, t)

	expected := "username-update"
	user := models.NewUser(expected)
	id := userMap["john.doe"].GetUserId()

	if &id == nil {
		t.Errorf("Is nil")
	}
	user.SetUserId(id)
	userRepository.Update(user)

	updatedUser, _ := userRepository.GetById(id)

	if updatedUser.GetUsername() != expected {
		t.Errorf("Updating of the username did not work, got %s, expected %s.", updatedUser.GetUsername(), expected)
	}
}
