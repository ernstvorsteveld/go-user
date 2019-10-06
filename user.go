package user

import (
	"github.com/google/uuid"
)

type User interface {
	GetUsername() UsernameType
	GetEmail() EmailType

	SetUserId() IdType
	GetUserId() IdType
}

func (u *UserIdentity) GetUsername() UsernameType {
	return u.user.username
}

func (u *UserIdentity) GetEmail() EmailType {
	return u.user.email
}

func (u *UserIdentity) SetUserId() IdType {
	var id IdType = IdType(uuid.New())
	u.user.id = id
	return id
}

func (u *UserIdentity) GetUserId() IdType {
	return u.user.id
}
