package handler

import (
	"example/gthomework/work7/api"
	"example/gthomework/work7/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

var ProdHandlerImpl ProdHandler

type ProdHandler struct {
	ProdServiceClient api.ProdServiceClient
}

func init() {
	// 1. 新建连接，端口是服务端开放的8082端口
	// 并且添加grpc.WithInsecure()，不然没有证书会报错
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	ProdHandlerImpl = ProdHandler{
		ProdServiceClient: api.NewProdServiceClient(conn),
	}
}

func (u *ProdHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":  "sucess",
		"data": "helloworld"})
}

func (u *ProdHandler) QueryProduct(c *gin.Context) {
	productIdStr := c.Query("productid")
	var err error
	var productId int64
	if productId, err = strconv.ParseInt(productIdStr, 10, 32); err != nil {
		utils.HandleErr(c, &err)
		return
	}

	if resp, err := u.ProdServiceClient.QueryProduct(c, &api.QueryProductRequest{ProductId: int32(productId)}); err != nil {
		utils.HandleErr(c, &err)
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0,
			"msg":  "sucess",
			"data": resp})
	}
	return
}
