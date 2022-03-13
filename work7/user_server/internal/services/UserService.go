package services

import (
	"context"
	"example/gthomework/work7/api"
	"example/gthomework/work7/user_server/internal/dal/dao"
	"example/gthomework/work7/user_server/internal/dal/dto"
	"fmt"
)

var UserServiceImpl *UserService

type UserService struct {
	UserDao *dao.UserDao
}

func init() {
	UserServiceImpl = &UserService{UserDao: dao.UserDaoImpl}
}

func (u *UserService) QueryUser(ctx context.Context, request *api.QueryUserRequest) (resp *api.QueryUserResponse, err error) {
	var userDto dto.User
	if userDto, err = u.UserDao.FindUserById(request.UserId); err != nil {
		return
	}
	fmt.Println("服务端处理请求QueryUser")
	resp = &api.QueryUserResponse{
		Code:   0,
		Msg:    "成功",
		UserId: userDto.ID,
		Name:   userDto.Name,
		Email:  userDto.Email,
		Age:    userDto.Age,
	}
	return
}
