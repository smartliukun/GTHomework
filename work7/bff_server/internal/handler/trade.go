package handler

import (
	"errors"
	"example/gthomework/work7/api"
	"example/gthomework/work7/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
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

	group, groupCtx := errgroup.WithContext(c)

	group.Go(func() error {
		if innerResp, err := UserHandlerImpl.userServiceClient.QueryUser(groupCtx, &api.QueryUserRequest{UserId: tradeReq.UserId}); err != nil {
			return err
		} else if innerResp.Code != 0 {
			return errors.New("查询账户信息异常,UserId=" + strconv.Itoa(int(tradeReq.UserId)))
		}
		return nil
	})

	group.Go(func() error {
		if innerResp, err := ProdHandlerImpl.ProdServiceClient.QueryProduct(groupCtx, &api.QueryProductRequest{ProductId: tradeReq.ProductId}); err != nil {
			return err
		} else if innerResp.Code != 0 {
			return errors.New("查询商品信息异常,UserId=" + strconv.Itoa(int(tradeReq.ProductId)))
		}
		return nil
	})

	if err := group.Wait(); err != nil {
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
