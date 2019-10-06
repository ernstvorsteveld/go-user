package user

import (
	"github.com/google/uuid"
)

type User interface {
	GetUsername() UsernameType
	SetUsername(username UsernameType)
	GetEmail() EmailType
	SetEmail(email EmailType)

	SetUserId() IdType
	GetUserId() IdType
}

func NewUser(username UsernameType) *UserIdentity {
	pUser := &UserIdentity{
		user: UserType{
			username: username,
		},
	}
	return pUser
}

func (u *UserIdentity) GetUsername() UsernameType {
	return u.user.username
}

func (u *UserIdentity) SetUsername(username UsernameType) {
	u.user.username = username
}

func (u *UserIdentity) GetEmail() EmailType {
	return u.user.email
}

func (u *UserIdentity) SetEmail(email EmailType) {
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
