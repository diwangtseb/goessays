package main

import (
	kError "github.com/go-kratos/kratos/v2/errors"
)

var ErrUserJoinRate = kError.NotFound("xx", "xxx")

func main() {
	ErrToKratosErr(ErrUserJoinRate)
}

func ErrToKratosErr(err error) *kError.Error {
	if err == nil {
		return nil
	}
	if kerr, ok := err.(*kError.Error); ok {
		return kerr
	} else {
		return kError.New(kError.UnknownCode, err.Error(), err.Error())
	}
}
