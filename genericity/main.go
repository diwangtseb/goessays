package main

import "fmt"

func Run[Any string](a Any) {
	fmt.Println(a)
}

func Add[T int | float32 | float64](a, b T) T {
	return a + b
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Animal interface {
	Dog | Cat
}

func Arun[T Animal](t T) {
	fmt.Println(t)
}

func main() {
	Arun(Dog{
		Name: "saner",
	})
}
