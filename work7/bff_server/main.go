package main

import (
	"example/gthomework/work7/bff_server/internal/handler"
	"example/gthomework/work7/bff_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//1.注册服务，通过wire自动装配
	registerHandler(r)
	r.Run(":8080")
}

func registerHandler(r *gin.Engine) {
	v1 := r.Group("/api")
	v1.Use(middleware.AopLog())
	v1.GET("/hello", handler.UserHandlerImpl.Hello)
	v1.GET("/user", handler.UserHandlerImpl.QueryUser)
	v1.GET("/product", handler.ProdHandlerImpl.QueryProduct)
	v1.POST("/trade", handler.TradeHandlerImpl.Trade)
}
