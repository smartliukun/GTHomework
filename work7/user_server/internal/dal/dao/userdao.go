package dao

import (
	"example/gthomework/work7/user_server/internal/dal/dto"
	"example/gthomework/work7/user_server/internal/dal/mysql"
	"fmt"
	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

var UserDaoImpl *UserDao

func init() {
	UserDaoImpl = &UserDao{
		DB: mysql.GetDB(),
	}

}

func (u *UserDao) FindUserById(userId int32) (user dto.User, err error) {
	if err = u.DB.Debug().Model(&dto.User{}).Where("id = ?", userId).Take(&user).Error; err != nil {
		return
	}
	fmt.Printf("get user =%v", user)
	return

}
