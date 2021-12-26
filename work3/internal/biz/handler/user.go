package handler

import (
	"example/gthomework/work3/api"
	"example/gthomework/work3/internal/biz/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (u *UserHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":  "sucess",
		"data": "helloworld"})
}

func (u *UserHandler) GetUserList(c *gin.Context) {
	inputName := c.Query("name")
	list, err := u.UserService.GetUserList(c, inputName)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 200,
			"msg":  "unknow error",
			"data": err.Error()})
		return
	}
	result := make([]api.UserVO, len(list))
	for i, entiy := range list {
		result[i].ID = entiy.ID
		result[i].Name = entiy.Name
		result[i].Age = entiy.Age
		result[i].Email = entiy.Email
		result[i].Birthday = entiy.Birthday
	}
	c.JSON(http.StatusOK, gin.H{"code": 0,
		"msg":  "sucess",
		"data": result})
}
