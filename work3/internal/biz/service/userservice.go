package service

import (
	"example/gthomework/work3/internal/biz/dal/dao"
	"example/gthomework/work3/internal/biz/dal/dto"
	"example/gthomework/work3/internal/biz/model"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{UserDao: userDao}
}

func (u *UserService) GetUserList(c *gin.Context, inputName string) (result []model.UserEntity, err error) {
	var listDto []dto.User
	if listDto, err = u.UserDao.FindUserList(inputName); err != nil {
		return
	}
	result = make([]model.UserEntity, len(listDto))
	for i, record := range listDto {
		result[i] = model.UserEntity{
			ID:       record.ID,
			Name:     record.Name,
			Email:    record.Email,
			Age:      record.Age,
			Birthday: record.Birthday,
		}
	}
	return
}
