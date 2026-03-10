package main

import "fmt"

type rect struct {
	width int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func Methods() {
	x := rect{width: 2, height: 4}
	fmt.Println(x)
	fmt.Println(x.area())
}