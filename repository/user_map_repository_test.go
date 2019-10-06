package repository

import (
	models "github.com/ernstvorsteveld/go-user"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func Test_create_and_retrieve_user(t *testing.T) {
	userRepository := NewMapUserRepository()

	pUser := &models.UserIdentity{
		user: models.UserType{
			id: models.IdType(uuid.New()),
			name: models.NameType{
				first: "John",
				last:  "Doe",
			},
		},
	}
	uuid, error := userRepository.Create(pUser)
	if error != nil {
		t.Errorf("Error returned by user repository while creating a user.")
	}
	if len(uuid) <= 0 {
		t.Errorf("When Creating a user in the repository, expected a UUID value.")
	}

	pUserFromGet, errorFromGet := userRepository.GetById(uuid)
	if errorFromGet != nil {
		t.Errorf("Error returned by user repository while getting a user.")
	}

	same := reflect.DeepEqual(*pUser, *pUserFromGet)
	if !same {
		t.Errorf("The retrieven user differs from the created user.")
	}
}
