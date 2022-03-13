package dto

type Trade struct {
	TradeNo    int32 `gorm:"primary_key"`
	UserId     int32
	ProductId  int32
	ProductNum int32
	Price      int32
	Cost       int64
	Date       string
}

func (Trade) TableName() string {
	return "trade"
}
