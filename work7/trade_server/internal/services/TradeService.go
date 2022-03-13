package services

import (
	"context"
	"example/gthomework/work7/api"
	"example/gthomework/work7/trade_server/internal/dal/dao"
	"example/gthomework/work7/trade_server/internal/dal/dto"
)

var TradeServiceImpl *TradeService

type TradeService struct {
	TradeDao *dao.TradeDao
}

func init() {
	TradeServiceImpl = &TradeService{TradeDao: dao.TradeDaoImpl}
}

func (t *TradeService) Trade(ctx context.Context, req *api.TradeRequest) (resp *api.TradeResponse, err error) {

	newTrade := dto.Trade{
		UserId:     req.UserId,
		ProductId:  req.ProductId,
		ProductNum: req.ProductNum,
		Price:      req.Price,
		Cost:       req.Cost,
		Date:       req.Date,
	}
	if err = t.TradeDao.CreateTrade(newTrade); err != nil {
		return
	}
	return &api.TradeResponse{
		Code:    0,
		Msg:     "成功",
		TradeNo: newTrade.TradeNo,
		Cost:    newTrade.Cost,
	}, nil
}
