package user

import (
  "time"
  "github.com/google/uuid"
  password "github.com/ernstvorsteveld/go-password"
)

type EmailType string

type LoginCodeType int

type NameType struct {
  first string
  middle string
  last string
}

const (
    UserName LoginCodeType = iota
    Email
)

type LoginType struct {
  loginType LoginCodeType
  value string
}

type LastLoginType struct {
  when time.Time
  what LoginCodeType
}

type UserType struct {
  id IdType
  name NameType
  email EmailType
  logins []LoginType
  created time.Time
  lastLogin LastLoginType
  password password.Password
}

type IdType uuid.UUID

type UserIdentity struct {
  user UserType
}
