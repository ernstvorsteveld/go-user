package repository

import (
  models "github.com/ernstvorsteveld/go-user"
  "github.com/google/uuid"
)

type userMapRepository struct {
  idMap map[uuid.UUID]*models.UserIdentity
  usernameMap map[string]*models.UserIdentity
  emailMap map[string]*models.UserIdentity
}

func NewMapUserRepository() *userMapRepository {
  return &userMapRepository{
    idMap : make(map[uuid.UUID]*models.UserIdentity),
    usernameMap : make(map[string]*models.UserIdentity),
    emailMap : make(map[string]*models.UserIdentity),
	}
}

func (m *userMapRepository) GetById(id uuid.UUID) (*models.UserIdentity, error) {
  return m.idMap[id], nil
}

func (m *userMapRepository) GetByLoginCode(loginCode string) (*models.UserIdentity, error) { 
    return nil, nil
}

func (m *userMapRepository) Update(user *models.UserIdentity) (*models.UserIdentity, error) {
    return nil, nil
}

func (m *userMapRepository) Create(u *models.UserIdentity) (uuid.UUID, error) {
  u.User.Id = uuid.New()
  m.idMap[u.User.Id] = u
  return u.User.Id, nil
}

func (m *userMapRepository) Delete(id uuid.UUID) (bool, error) {
    return false, nil
}
