package repository

import (
	models "github.com/ernstvorsteveld/go-user"
	"reflect"
	"testing"
)

func Test_create_and_retrieve_user(t *testing.T) {
	userRepository := NewMapUserRepository()
	userMap := createUsers(userRepository, t)

	john := userMap["john.doe"]
	pJohn := &john
	pUserFromGet, errorFromGet := userRepository.GetById(pJohn.GetUserId())
	if errorFromGet != nil {
		t.Errorf("Error returned by user repository while getting a user.")
	}

	same := reflect.DeepEqual(userMap["john.doe"], *pUserFromGet)
	if !same {
		t.Errorf("The user retrieved %s differs from the created user %s.", pJohn.GetUsername(), pUserFromGet.GetUsername())
	}

	pUsernameUser, error := userRepository.GetByLoginCode("john.doe")
	if error != nil {
		t.Errorf("Error occurred: %s", error)
	}
	if pUsernameUser != pJohn {
		t.Errorf("Users are not the same.")
	}
}

func createUsers(userRepository UserRepository, t *testing.T) map[string]models.UserIdentity {
	userMap := make(map[string]*models.UserIdentity)

	username := "john.doe"
	pUser := models.NewUser(username)
	uuid, error := userRepository.Create(pUser)
	validateUserCreate(error, uuid, t)
	userMap[username] = pUser

	username = "sander.kuiper"
	pUser = models.NewUser(username)
	uuid, error = userRepository.Create(pUser)
	validateUserCreate(error, uuid, t)
	userMap[username] = pUser

	username = "annelies.gerritsen"
	pUser = models.NewUser(username)
	uuid, error = userRepository.Create(pUser)
	validateUserCreate(error, uuid, t)
	userMap[username] = pUser

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
