package repository

import (
	models "github.com/ernstvorsteveld/go-user"
)

type UserRepository interface {
	GetById(id models.IdType) (models.User, error)
	GetByLoginCode(loginCode string) (*models.User, error)
	Update(u models.User) (*models.User, error)
	Create(u models.User) (*models.IdType, error)
	Delete(id *models.IdType) (bool, error)
}
