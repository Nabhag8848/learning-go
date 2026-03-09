package main

import "fmt"

func Pointers(){
	var x = 10

	passByValue(x)
	fmt.Println(x)
	passByReference(&x)
	fmt.Println(x)
	fmt.Println(&x)
}

func passByValue(i int) {
	i = 0;
}

func passByReference(i *int) {
	*i = 0
}

