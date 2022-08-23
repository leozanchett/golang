package main

import "fmt"

type Animal interface {
	Roar() string
	Run() string
}

func RoarAndRun(a Animal) {
	fmt.Printf("%s and %s\n", a.Roar(), a.Run())
}

type Dog struct{}

func (d Dog) Roar() string {
	return "Woof!"
}

func (d Dog) Run() string {
	return "Run like a dog"
}

type Cat struct{}

func (c *Cat) Roar() string {
	return "Meow!"
}

func (c *Cat) Run() string {
	return "Run like a cat"
}

func main() {
	myDog := Dog{}
	myCat := Cat{}

	RoarAndRun(myDog)
	RoarAndRun(&myCat)
}
