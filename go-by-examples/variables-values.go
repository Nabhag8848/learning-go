package main

import (
	"fmt"
	"math"
)

const VALUE string = "constant"

func Basics() {
	fmt.Println(1.0 / 3.0)
	fmt.Println((math.Pow(3, 2)))

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	e := "declare&initialize"
	fmt.Println(e)

	var f int
	fmt.Println(f)
	
	var g bool
	fmt.Println(g)
	
	var h any
	fmt.Println(h)

	h = "nil to string"

	fmt.Println(h)

	const x int = 10
	// x = 23 not allowed

}