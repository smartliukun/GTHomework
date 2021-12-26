//+build wireinject

package main

import (
	"example/gthomework/work3/internal/biz/dal/dao"
	"example/gthomework/work3/internal/biz/dal/mysql"
	"example/gthomework/work3/internal/biz/handler"
	"example/gthomework/work3/internal/biz/service"
	"github.com/google/wire"
)

func InitializeEvent() (*handler.UserHandler, error) {
	wire.Build(mysql.GetDB, dao.NewUserDao, service.NewUserService, handler.NewUserHandler)
	return &handler.UserHandler{}, nil
}
