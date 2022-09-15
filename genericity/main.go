package main

import "fmt"

type Zoo[T any] struct {
	Animals []*T
}

func (z *Zoo[T]) Run() {
	for _, a := range z.Animals {
		fmt.Println(a, "is ok!")
	}
}

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
	z := Zoo[Dog]{}
	z.Animals = append(z.Animals, &Dog{
		Name: "x3",
	})
	z.Animals = append(z.Animals, &Dog{
		Name: "x4",
	})
	z.Run()
}
