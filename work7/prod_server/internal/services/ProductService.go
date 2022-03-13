package services

import (
	"context"
	"example/gthomework/work7/api"
	"example/gthomework/work7/prod_server/internal/dal/dao"
	"example/gthomework/work7/prod_server/internal/dal/dto"
	"fmt"
)

var ProdServiceImpl *ProdService

type ProdService struct {
	ProdDao *dao.ProdDao
}

func init() {
	ProdServiceImpl = &ProdService{ProdDao: dao.ProdDaoImpl}
}

func (u *ProdService) QueryProduct(ctx context.Context, request *api.QueryProductRequest) (resp *api.QueryProductResponse, err error) {
	fmt.Println("服务端处理请求")

	var productDto dto.Product
	if productDto, err = u.ProdDao.FindProductById(request.ProductId); err != nil {
		return
	}
	fmt.Println("服务端处理请求QueryProduct")
	resp = &api.QueryProductResponse{
		Code:      0,
		Msg:       "成功",
		ProductId: productDto.ProductID,
		Name:      productDto.Name,
		Price:     productDto.Price,
		Stock:     productDto.Stock,
	}
	return
}
