package repository

import (
	"errors"
	models "github.com/ernstvorsteveld/go-user"
)

type userMapRepository struct {
	idMap       map[models.IdType]*models.UserIdentity
	usernameMap map[string]*models.UserIdentity
	emailMap    map[string]*models.UserIdentity
}

func NewMapUserRepository() *userMapRepository {
	return &userMapRepository{
		idMap:       make(map[models.IdType]*models.UserIdentity),
		usernameMap: make(map[string]*models.UserIdentity),
		emailMap:    make(map[string]*models.UserIdentity),
	}
}

func (m *userMapRepository) GetById(id models.IdType) (*models.UserIdentity, error) {
	return m.idMap[id], nil
}

func (m *userMapRepository) GetByLoginCode(loginCode string) (*models.UserIdentity, error) {
	pUser := m.usernameMap[loginCode]
	if pUser == nil {
		return nil, errors.New("Could not find user with loginCode: " + loginCode)
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
