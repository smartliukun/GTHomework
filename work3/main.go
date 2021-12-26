package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type ExampleApp struct {
	db *sql.DB
}

var exampleApp ExampleApp

func main() {
	r := gin.Default()
	//1.注册服务，通过wire自动装配
	{
		userHandler, err := InitializeEvent()
		if err != nil {
			panic(err)
		}
		registerHandler(r, userHandler)
	}

	//2.异步启动服务
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}
	go func() {
		if curErr := srv.ListenAndServe(); curErr != nil && curErr != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", curErr)
		}
	}()
	//3.监听中断
	// 一个通知退出的chan，接收中断信息，并关闭服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sign, ok := <-quit
	log.Println(fmt.Sprintf("sign=%v,ok=%v", sign, ok))
	srv.Shutdown(context.Background())

}
