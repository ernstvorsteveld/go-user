package repository

import (
	"errors"
	"fmt"
	models "github.com/ernstvorsteveld/go-user"
)

type userMapRepository struct {
	idMap       map[models.IdType]models.User
	usernameMap map[string]models.User
	emailMap    map[string]models.User
}

func NewMapUserRepository() *userMapRepository {
	return &userMapRepository{
		idMap:       make(map[models.IdType]models.User),
		usernameMap: make(map[string]models.User),
		emailMap:    make(map[string]models.User),
	}
}

func (m *userMapRepository) GetById(id models.IdType) (models.User, error) {
	pUser := m.idMap[id]
	if &pUser == nil {
		return nil, errors.New(fmt.Sprintf("User with id %s does not exist.", id))
	}
	return pUser, nil
}

func (m *userMapRepository) GetByLoginCode(loginCode string) (*models.User, error) {
	pUser := m.usernameMap[loginCode]
	if &pUser == nil {
		return nil, errors.New(fmt.Sprintf("Could not find user with loginCode %s.", loginCode))
	}
	return &pUser, nil
}

func (m *userMapRepository) Update(userIdenity models.User) (*models.User, error) {
	return nil, nil
}

func (m *userMapRepository) Create(u models.User) (*models.IdType, error) {
	id := u.SetUserId()
	m.idMap[id] = u
	m.usernameMap[u.GetUsername()] = u
	return &id, nil
}

func (m *userMapRepository) Delete(id models.IdType) (bool, error) {
	return false, nil
}
