package main

import "fmt"

type Rectangle struct {
	Height, Width int
}

func (r Rectangle) Surface() int {
	return r.Height * r.Width
}

func (r Rectangle) SemPonteiro(factor int) {
	r.Height *= factor
	r.Width *= factor
}

func (r *Rectangle) ComPonteiro(factor int) {
	r.Height *= factor
	r.Width *= factor
}

func main() {
	r := Rectangle{Height: 2, Width: 3}
	fmt.Printf("rectangle %v has surface %d\n", r, r.Surface())
	r.SemPonteiro(2)
	fmt.Printf("rectangle %v has surface %d\n", r, r.Surface())
	r.ComPonteiro(2)
	fmt.Printf("rectangle %v has surface %d\n", r, r.Surface())
}
