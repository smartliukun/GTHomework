package handler

import (
	"example/gthomework/work7/api"
	"example/gthomework/work7/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var TradeHandlerImpl TradeHandler

type TradeHandler struct {
	TradeServiceClient api.TradeServiceClient
}

func init() {
	// 1. 新建连接，端口是服务端开放的8082端口
	// 并且添加grpc.WithInsecure()，不然没有证书会报错
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	TradeHandlerImpl = TradeHandler{
		TradeServiceClient: api.NewTradeServiceClient(conn),
	}
}

func (u *TradeHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":  "sucess",
		"data": "helloworld"})
}

type TradeReq struct {
	UserId     int32
	ProductId  int32
	ProductNum int32
	Price      int32
	Date       string
}

func (u *TradeHandler) Trade(c *gin.Context) {
	var tradeReq TradeReq
	if err := c.Bind(&tradeReq); err != nil {
		utils.HandleErr(c, &err)
		return
	}

	if resp, err := u.TradeServiceClient.Trade(c, &api.TradeRequest{
		UserId:     tradeReq.UserId,
		ProductId:  tradeReq.ProductId,
		ProductNum: tradeReq.ProductNum,
		Price:      tradeReq.Price,
		Cost:       int64(tradeReq.Price * tradeReq.ProductNum),
		Date:       tradeReq.Date,
	}); err != nil {
		utils.HandleErr(c, &err)
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0,
			"msg":  "sucess",
			"data": resp})
	}
	return
}
