package main

import "errors"

func main() {
	client := NewSdk("asd", "sdd")
	err := client.Do()
	if err != nil {
		panic(err)
	}
}

type SDKer interface {
	Do() error
}

type indicator func() error

type sdk struct {
	auth       func() (bool, error)
	middleware []func() error
	handle     indicator
}

func NewSdk(ak string, as string) SDKer {
	sdk := &sdk{
		auth: func() (bool, error) {
			if ak != "" || as != "" {
				return false, errors.New("auth failed")
			}
			return true, nil
		},
	}
	defer sdk.initMiddleware()
	return sdk
}

func (sdk *sdk) initMiddleware() {
	auth := func() error {
		_, err := sdk.auth()
		if err != nil {
			return err
		}
		return nil
	}
	sdk.middleware = append(sdk.middleware, auth)
	sdk.handle = func() error {
		for _, v := range sdk.middleware {
			if v() != nil {
				return v()
			}
		}
		return nil
	}
}

func (s *sdk) Do() error {
	return s.handle()
}
