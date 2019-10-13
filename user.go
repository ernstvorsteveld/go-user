package user

import (
	"github.com/google/uuid"
)

type IdType uuid.UUID

type User interface {
	GetUsername() string
	SetUsername(username string)
	GetEmail() string
	SetEmail(email string)

	InitId() IdType
	SetUserId(id IdType) IdType
	GetUserId() IdType
}
