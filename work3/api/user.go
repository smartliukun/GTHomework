package api

import "time"

type UserVO struct {
	ID       int32
	Name     string
	Email    string
	Age      int32
	Birthday time.Time
}
