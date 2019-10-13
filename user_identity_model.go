package user

import (
	password "github.com/ernstvorsteveld/go-password"
	"time"
)

const (
	UserName int = iota
	Email
)

type NameType struct {
	first  string
	middle string
	last   string
}

type LoginType struct {
	loginType int
	value     string
}

type LastLoginType struct {
	when time.Time
	what int
}

type UserType struct {
	username  string
	name      NameType
	email     string
	logins    []LoginType
	created   time.Time
	lastLogin LastLoginType
	password  password.Password
}

type UserIdentity struct {
	id   IdType   `json:"id"`
	user UserType `json:"user"`
}
