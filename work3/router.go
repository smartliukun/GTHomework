package main

import (
	"example/gthomework/work3/internal/biz/handler"
	"example/gthomework/work3/internal/biz/middleware"
	"github.com/gin-gonic/gin"
)

func registerHandler(r *gin.Engine, userHandler *handler.UserHandler) {
	v1 := r.Group("/api")
	v1.Use(middleware.AopLog())
	v1.GET("/hello", userHandler.Hello)
	v1.GET("/userlist", userHandler.GetUserList)
}
