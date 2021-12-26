package model

import "time"

type UserEntity struct {
	ID       int32
	Name     string
	Email    string
	Age      int32
	Birthday time.Time
}
