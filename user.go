package user

import (
	"github.com/google/uuid"
)

type User interface {
	GetUsername() string
	SetUsername(username string)
	GetEmail() string
	SetEmail(email string)

	SetUserId() IdType
	GetUserId() IdType
}

func NewUser(username string) *UserIdentity {
	pUser := &UserIdentity{
		user: UserType{
			username: username,
		},
	}
	return pUser
}

func (u *UserIdentity) GetUsername() string {
	return u.user.username
}

func (u *UserIdentity) SetUsername(username string) {
	u.user.username = username
}

func (u *UserIdentity) GetEmail() string {
	return u.user.email
}

func (u *UserIdentity) SetEmail(email string) {
	u.user.email = email
}

func (u *UserIdentity) SetUserId() IdType {
	var id IdType = IdType(uuid.New())
	u.id = id
	return id
}

func (u *UserIdentity) GetUserId() IdType {
	return u.id
}
