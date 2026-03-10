package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return "describe"
}

type container struct {
	base
	name string
}

func Embeddings() {
	c := container{base: base{num: 1},name: "nabhag" }
	var a describer = c
	a.describe()
	fmt.Println(c.base.num, c.name, c.num, c.base.describe(), c.describe(), a.describe())


}

type describer interface {
	describe() string
}
