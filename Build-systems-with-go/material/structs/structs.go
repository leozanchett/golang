package main

import "fmt"

type Retangle struct {
	width, height int
}

func NewRetangle(width, height int) *Retangle {
	return &Retangle{width, height}
}

func NovoRetangulo(width, height int) Retangle {
	return Retangle{width, height}
}

func main() {
	a := Retangle{}
	fmt.Println(a)

	b := Retangle{width: 4, height: 4}
	fmt.Println(b)

	c := Retangle{10, 3}
	fmt.Println(c)

	d := Retangle{width: 10}
	fmt.Println(d)

	e := NewRetangle(2, 3)
	fmt.Println(e)
	fmt.Println(*e)

	f := NovoRetangulo(3, 4)
	fmt.Println(f)
}
