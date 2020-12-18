package model

import (
	"HelloWorld/api"
	"time"
)

type User struct {
	Id    int64
	Name  string
	Mtime time.Time
	Ctime time.Time
}

func (u User) ToReply() *api.UserReply {
	return &api.UserReply{
		Id:   u.Id,
		Name: u.Name,
	}
}
