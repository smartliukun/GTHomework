package dao

import (
	"example/gthomework/work7/trade_server/internal/dal/dto"
	"example/gthomework/work7/trade_server/internal/dal/mysql"
	"fmt"
	"gorm.io/gorm"
)

type TradeDao struct {
	DB *gorm.DB
}

var TradeDaoImpl *TradeDao

func init() {
	TradeDaoImpl = &TradeDao{
		DB: mysql.GetDB(),
	}

}

func (u *TradeDao) FindTrade(tradeNo int32) (trade dto.Trade, err error) {
	if err = u.DB.Model(&dto.Trade{}).Where("trade_no = ?", tradeNo).Take(&trade).Error; err != nil {
		return
	}
	fmt.Printf("get trade =%v", trade)
	return

}

func (u *TradeDao) CreateTrade(newTrade dto.Trade) (err error) {
	if err = u.DB.Model(&dto.Trade{}).Create(newTrade).Error; err != nil {
		return
	}
	fmt.Printf("CreateTrade =%v", newTrade)
	return

}
