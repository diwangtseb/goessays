package main

import "fmt"

func main() {
	var input string
	fmt.Println("your images are ready to be pushed,yes or no?  ")
	fmt.Scanln(&input)
	if input != "yes" {
		fmt.Println("bye")
		return
	}
	var inputUserName string
	fmt.Println("please input your username:")
	fmt.Scanf("%s:\n", &inputUserName)
	var inputPwd string
	fmt.Println("please input your password:")
	fmt.Scanf("%s:\n", &inputPwd)
	if inputUserName == "" || inputPwd == "" {
		fmt.Println("username or password is empty")
		return
	}
}
