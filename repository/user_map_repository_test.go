package repository

import (
  "testing"
  models "github.com/ernstvorsteveld/go-user"
  "github.com/google/uuid"
)

func Test_implement_the_repository(t *testing.T) {
  userRepository := NewMapUserRepository()
  
  pUser := &models.UserIdentity{
    User: models.UserType{
      Id: uuid.New(),
      Name: models.NameType{
        First: "John",
        Last: "Doe",
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
}