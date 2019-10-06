package repository

import (
	"errors"
	models "github.com/ernstvorsteveld/go-user"
)

type userMapRepository struct {
	idMap       map[models.IdType]*models.UserIdentity
	usernameMap map[models.UsernameType]*models.UserIdentity
	emailMap    map[models.EmailType]*models.UserIdentity
}

func NewMapUserRepository() *userMapRepository {
	return &userMapRepository{
		idMap:       make(map[models.IdType]*models.UserIdentity),
		usernameMap: make(map[models.UsernameType]*models.UserIdentity),
		emailMap:    make(map[models.EmailType]*models.UserIdentity),
	}
}

func (m *userMapRepository) GetById(id models.IdType) (*models.UserIdentity, error) {
	return m.idMap[id], nil
}

func (m *userMapRepository) GetByLoginCode(loginCode models.UsernameType) (*models.UserIdentity, error) {
	pUser := m.usernameMap[models.UsernameType(loginCode)]
	if pUser == nil {
		return nil, errors.New("Could not find user with loginCode: " + string(loginCode))
	}
	return pUser, nil
}

func (m *userMapRepository) Update(user *models.UserIdentity) (*models.UserIdentity, error) {
	return nil, nil
}

func (m *userMapRepository) Create(u *models.UserIdentity) (models.IdType, error) {
	m.idMap[u.SetUserId()] = u
	m.usernameMap[u.GetUsername()] = u
	return u.GetUserId(), nil
}

func (m *userMapRepository) Delete(id models.IdType) (bool, error) {
	return false, nil
}
