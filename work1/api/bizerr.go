package api

import "fmt"

type BizError struct {
	Ret RetEnm
	Err error
}

func (bizErr BizError) Error() string {
	return fmt.Sprintf("BizError=%v,err=%v", bizErr.Ret, bizErr.Err)
}

func (bizErr *BizError) Unwrap() error { return bizErr.Err }
