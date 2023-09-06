package main

import (
	"context"
	"fmt"

	"github.com/superyyk/rpc_models/db"
	model "github.com/superyyk/rpc_models/mdel"
)

type Req struct {
	Id string
	Ty int
}
type Resp struct {
	UserInfo model.UsersInfo
}

type User struct {
}

var Db = db.Db

func (u *User) GetUserByID(ctx context.Context, req *Req, reply *Resp) {
	fmt.Println("hello")
	var user model.UsersInfo
	Db.Table("users").Where("uid=?", req.Id).First(&user)
	reply.UserInfo = user
}
