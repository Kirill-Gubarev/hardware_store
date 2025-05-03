package db

import (
	"time"
)

type Role struct{
	Id    *string  `json:"id"`
	Name  *string  `json:"name"`
}

type User struct{
	Id        *string  `json:"id"`
	Login     *string  `json:"login"`
	Password  *string  `json:"password"`
	Role      *Role    `json:"role"`
}

type Session struct{
	Id         *string     `json:"id"`
	User       *User       `json:"user"`
	Token      *string     `json:"token"`
	CreatedAt  *time.Time  `json:"createdAt"`
}

func NewRole()(Role){
	return Role{}
}
func NewUser()(User){
	return User{Role: &Role{}}
}
func NewSession()(Session){
	return Session{User: &User{Role: &Role{}}}
}
