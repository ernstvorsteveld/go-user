package user

import (
  "time"
  "github.com/google/uuid"
  "github.com/ernstvorsteveld/go-password"
)

type EmailType string

type LoginCodeType int

type NameType struct {
  first string    `bson:"first_name" json:"first_name"`
  middle string   `bson:"middle_name" json:"middle_name"`
  last string     `bson:"last_name" json:"last_name"`
}

const (
    UserName LoginCodeType = iota
    Email
)

type LoginType struct {
  loginType LoginCodeType `bson:"login_type" json:"login_type"`
  value string            `bson:"value" json:"value"`
}

type LastLoginType struct {
  when Time       `bson:"when" json:"when"`
  what LoginType  `bson:"what" json:"what"`
}

type User struct {
  name NameType             `bson:"name" json:"name"`
  email EmailType           `bson:"email" json:"email"`
  logins []LoginType        `bson:"login_codes" json:"login_codes"`
  created Time              `bson:"create_time" json:"create_time"`
  lastLogin LastLoginType   `bson:"last_login" json:"last_login"`
  password BcryptPassword   `bson:"password" json:"password"`
}

type userIdentity struct {
  ID  bson.ObjectId   `bson:"_id"`
  id UUID             `bson:"id" json:"id"`
  user User           `bson:"user" json:"user"`
}