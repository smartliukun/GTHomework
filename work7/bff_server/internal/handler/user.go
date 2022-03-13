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

var UserHandlerImpl UserHandler

type UserHandler struct {
	userServiceClient api.UserServiceClient
}

func init() {
	// 1. 新建连接，端口是服务端开放的8082端口
	// 并且添加grpc.WithInsecure()，不然没有证书会报错
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	UserHandlerImpl = UserHandler{
		userServiceClient: api.NewUserServiceClient(conn),
	}
}

func (u *UserHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":  "sucess",
		"data": "helloworld"})
}

func (u *UserHandler) QueryUser(c *gin.Context) {
	inputUserIdStr := c.Query("userid")
	var err error
	var inputUserId int64
	if inputUserId, err = strconv.ParseInt(inputUserIdStr, 10, 32); err != nil {
		utils.HandleErr(c, &err)
		return
	}

	if resp, err := u.userServiceClient.QueryUser(c, &api.QueryUserRequest{UserId: int32(inputUserId)}); err != nil {
		utils.HandleErr(c, &err)
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0,
			"msg":  "sucess",
			"data": resp})
	}
	return
}
