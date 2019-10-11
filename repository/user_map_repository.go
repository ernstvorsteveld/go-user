package repository

import (
	"errors"
	"fmt"
	models "github.com/ernstvorsteveld/go-user"
)

type userMapRepository struct {
	idMap       map[models.IdType]models.UserIdentity
	usernameMap map[string]models.UserIdentity
	emailMap    map[string]models.UserIdentity
}

func NewMapUserRepository() *userMapRepository {
	return &userMapRepository{
		idMap:       make(map[models.IdType]models.UserIdentity),
		usernameMap: make(map[string]models.UserIdentity),
		emailMap:    make(map[string]models.UserIdentity),
	}
}

func (m *userMapRepository) GetById(id models.IdType) (*models.UserIdentity, error) {
	pUser := m.idMap[id]
	if &pUser == nil {
		return nil, errors.New(fmt.Sprintf("User with id %s does not exist.", id))
	}
	return &pUser, nil
}

func (m *userMapRepository) GetByLoginCode(loginCode string) (*models.UserIdentity, error) {
	pUser := m.usernameMap[loginCode]
	if &pUser == nil {
		return nil, errors.New(fmt.Sprintf("Could not find user with loginCode %s.", loginCode))
	}
	return &pUser, nil
}

func (m *userMapRepository) Update(userIdenity models.UserIdentity) (*models.UserIdentity, error) {
	return nil, nil
}

func (m *userMapRepository) Create(u models.UserIdentity) (*models.IdType, error) {
	id := u.SetUserId()
	m.idMap[id] = u
	m.usernameMap[u.GetUsername()] = u
	return &id, nil
}

func (m *userMapRepository) Delete(id models.IdType) (bool, error) {
	return false, nil
}
