package dao

import (
	"example/gthomework/work3/internal/biz/dal/dto"
	"fmt"
	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{DB: db}
}

func (u *UserDao) FindUserList(inputName string) (list []dto.User, err error) {
	if err = u.DB.Where("name like ?", "%"+inputName+"%").Find(&list).Error; err != nil {
		return
	}
	fmt.Printf("fetch user list=%v", list)
	return

}
