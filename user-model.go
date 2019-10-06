package user

import (
	password "github.com/ernstvorsteveld/go-password"
	"github.com/google/uuid"
	"time"
)

type EmailType string
type LoginCodeType int
type UsernameType string
type IdType uuid.UUID

const (
	UserName LoginCodeType = iota
	Email
)

type NameType struct {
	first  string
	middle string
	last   string
}

type LoginType struct {
	loginType LoginCodeType
	value     string
}

type LastLoginType struct {
	when time.Time
	what LoginCodeType
}

type UserType struct {
	username  UsernameType
	name      NameType
	email     EmailType
	logins    []LoginType
	created   time.Time
	lastLogin LastLoginType
	password  password.Password
}

type UserIdentity struct {
	id   IdType   `json:"id"`
	user UserType `json:"user"`
}
