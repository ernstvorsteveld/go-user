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

func (m *userMapRepository) GetByLoginCode(loginCode string) (models.User, error) {
	if loginCode == "" {
		return nil, errors.New("No loginCode provided.")
	}
	user := m.usernameMap[loginCode]
	if &user == nil {
		return nil, errors.New(fmt.Sprintf("Could not find user with loginCode %s.", loginCode))
	}
	return user, nil
}

func (m *userMapRepository) Update(u models.User) (models.User, error) {
	pId := u.GetUserId()
	if &pId == nil {
		return nil, errors.New(fmt.Sprintf("Cannot update a user with id == nil."))
	}
	user := m.idMap[pId]
	if &user == nil {
		return nil, errors.New(fmt.Sprintf("Could not find user to updete with id %s.", u.GetUserId()))
	}
	m.idMap[pId] = u
	m.usernameMap[u.GetUsername()] = u

	return user, nil
}

func (m *userMapRepository) Create(u models.User) (models.IdType, error) {
	id := u.InitId()
	m.idMap[id] = u
	m.usernameMap[u.GetUsername()] = u
	return id, nil
}

func (m *userMapRepository) Delete(id models.IdType) (bool, error) {
	return false, nil
}
