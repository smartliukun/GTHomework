package dto

type Product struct {
	ProductID int32 `gorm:"primary_key"`
	Name      string
	Price     int64
	Stock     int32
}

func (Product) TableName() string {
	return "product"
}
