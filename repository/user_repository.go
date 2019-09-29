package repository

import (
  models "github.com/ernstvorsteveld/go-user"
  "github.com/google/uuid"
)

type UserRepository interface {
//   Fetch(cursor string, num int64) ([]*models.Article, error)
	GetById(id uuid.UUID) (*models.UserIdentity, error)
	GetByLoginCode(loginCode string) (*models.UserIdentity, error)
	Update(user *models.UserIdentity) (*models.UserIdentity, error)
	Create(a *models.UserIdentity) (uuid.UUID, error)
	Delete(id uuid.UUID) (bool, error)
}