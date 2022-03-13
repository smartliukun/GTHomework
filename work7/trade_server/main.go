package main

import (
	tradeservices "example/gthomework/work7/api"
	"example/gthomework/work7/trade_server/internal/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 1. new一个grpc的server
	rpcServer := grpc.NewServer()

	// 2. 将刚刚我们新建的ProdService注册进去
	tradeservices.RegisterTradeServiceServer(rpcServer, services.TradeServiceImpl)

	// 3. 新建一个listener，以tcp方式监听8082端口
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	// 4. 运行rpcServer，传入listener
	_ = rpcServer.Serve(listener)
}