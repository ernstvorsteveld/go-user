package repository

import (
	models "github.com/ernstvorsteveld/go-user"
)

type UserRepository interface {
	GetById(id models.IdType) (*models.UserIdentity, error)
	GetByLoginCode(loginCode string) (*models.UserIdentity, error)
	Update(u models.UserIdentity) (*models.UserIdentity, error)
	Create(u models.UserIdentity) (*models.IdType, error)
	Delete(id models.IdType) (bool, error)
}
