package user

import (
	password "github.com/ernstvorsteveld/go-password"
	"github.com/google/uuid"
	"time"
)

type EmailType string

type LoginCodeType int

type NameType struct {
	First  string
	Middle string
	Last   string
}

const (
	UserName LoginCodeType = iota
	Email
)

type LoginType struct {
	loginType LoginCodeType
	value     string
}

type LastLoginType struct {
	when time.Time
	what LoginCodeType
}

type UsernameType string

type UserType struct {
	Id        uuid.UUID
	Username  UsernameType
	Name      NameType
	Email     EmailType
	logins    []LoginType
	created   time.Time
	lastLogin LastLoginType
	password  password.Password
}

type IdType uuid.UUID

type UserIdentity struct {
	User UserType
}

func GetUsername(u *UserIdentity) UsernameType {
	return u.User.Username
}

func GetEmail(u *UserIdentity) EmailType {
	return u.User.Email
}
