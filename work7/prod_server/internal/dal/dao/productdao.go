package dao

import (
	"example/gthomework/work7/prod_server/internal/dal/dto"
	"example/gthomework/work7/prod_server/internal/dal/mysql"
	"fmt"
	"gorm.io/gorm"
)

type ProdDao struct {
	DB *gorm.DB
}

var ProdDaoImpl *ProdDao

func init() {
	ProdDaoImpl = &ProdDao{
		DB: mysql.GetDB(),
	}

}

func (u *ProdDao) FindProductById(productId int32) (product dto.Product, err error) {
	if err = u.DB.Model(&dto.Product{}).Where("product_id = ?", productId).Take(&product).Error; err != nil {
		return
	}
	fmt.Printf("get user =%v", product)
	return

}
