package repository

import (
	models "github.com/ernstvorsteveld/go-user"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetById(id uuid.UUID) (*models.UserIdentity, error)
	GetByLoginCode(loginCode string) (*models.UserIdentity, error)
	Update(user *models.UserIdentity) (*models.UserIdentity, error)
	Create(a *models.UserIdentity) (models.IdType, error)
	Delete(id models.IdType) (bool, error)
}
