package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 22
	return &p
}

func Struct() {

	a := person{"nabhag", 22}
	b := person{name: "nabhag", age: 22}
	d := &person{name: "nabhag", age: 22}
	c := person{age: 22}
	e := d
	f := b
	fmt.Println(newPerson("nabhag"))

	fmt.Println(a,b,c,d,e,f)
	
	dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }

	fmt.Println(dog)
}