package main

import (
	"fmt"
	"strconv"
)

type User struct {
	ID   int
	Name string
}

func main() {
	users := make([]*User, 0)
	for i := 0; i < 10; i++ {
		user := &User{
			ID:   i,
			Name: "user" + strconv.Itoa(i),
		}
		users = append(users, user)
	}
	aUsers := make([]*User, 0)
	for _, user := range users {
		if user.ID%2 != 0 {
			aUsers = append(aUsers, user)
		}
	}
	fmt.Println(aUsers)
}
