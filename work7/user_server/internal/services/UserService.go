package services

import (
	"context"
	"example/gthomework/work7/api"
	"example/gthomework/work7/user_server/internal/dal/dao"
	"example/gthomework/work7/user_server/internal/dal/dto"
	"example/gthomework/work7/user_server/internal/redisservice"
	"fmt"
	"golang.org/x/sync/singleflight"
	"strconv"
	"time"
)

var UserServiceImpl *UserService

type UserService struct {
	UserDao *dao.UserDao
	gsf     singleflight.Group
}

func init() {
	UserServiceImpl = &UserService{UserDao: dao.UserDaoImpl}
}

func (u *UserService) QueryUser(ctx context.Context, request *api.QueryUserRequest) (resp *api.QueryUserResponse, err error) {

	fmt.Println("服务端处理请求QueryUser")
	key := "[user]_" + strconv.Itoa(int(request.UserId))

	var userDto dto.User
	if redisErr := redisservice.RedisServiceImpl.GetCache(ctx, key, &userDto); redisErr == nil {
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

	//如果缓存没有值，则通过singleFly，本地加锁获取数据库记录
	u.gsf.Do(key, func() (interface{}, error) {
		userDto, err = u.UserDao.FindUserById(request.UserId)
		return userDto, err
	})

	if err != nil {
		return nil, err
	}

	resp = &api.QueryUserResponse{
		Code:   0,
		Msg:    "成功",
		UserId: userDto.ID,
		Name:   userDto.Name,
		Email:  userDto.Email,
		Age:    userDto.Age,
	}

	var expire time.Duration
	expire = time.Duration(100) * time.Second
	redisservice.RedisServiceImpl.SetCache(ctx, key, expire, &userDto)

	return
}
