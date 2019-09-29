package repository

import (
  "testing"
)

func Test_implement_the_repository(t *testing.T) {
  userRepository := NewMapUserRepository()
  
  userRepository.Create(nil)
}