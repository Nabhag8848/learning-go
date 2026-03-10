package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}


func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) (float64, float64){
	return g.area(), g.perim()
}

func Interface() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	fmt.Println(measure(r))
	fmt.Println(measure(c))
}