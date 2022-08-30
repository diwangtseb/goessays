package main

import (
	"fmt"
)

type SelfError struct {
	errText string
}

func New(text string) *SelfError {
	return &SelfError{errText: text}
}

func (se *SelfError) Error() string {
	return se.errText
}

func main() {
	if err := New("xxx"); err != nil {
		fmt.Println(err.Error())
	}
}
