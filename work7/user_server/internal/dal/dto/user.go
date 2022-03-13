package dto

import (
	"time"
)

type User struct {
	ID       int32 `gorm:"primary_key"`
	Name     string
	Email    string
	Age      int32
	Birthday time.Time
}
