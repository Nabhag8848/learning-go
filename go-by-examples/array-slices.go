package main

import (
	"fmt"
	"slices"
)

func ArrayDS() {

	var a [5]int
	fmt.Println(a)

	b := []int{1,2,3,4}
	fmt.Println(b)

	c := [...]int{1,2,3,4}
	fmt.Println(c)

	d := [...]int{1, 2: 400,4}
	fmt.Println(d)

	twodArray := [2][3]int {
		{1,2,3}, 
		{1,2,3},
	}

	fmt.Println(twodArray)

	var arr [2][3]int

	for i := range 2 {
		for j := range 3 {
			arr[i][j] = i * j
		}
	}

	fmt.Println(arr)


	var slice []int
	fmt.Println(slice, len(slice))

	slice = make([]int, 3)
	fmt.Println(slice)

	slice = append(slice, 4)
	fmt.Println(slice)

	var slice1 []int = make([]int, len(slice))
	copy(slice1, slice) // values are copied not made pointer
	fmt.Println(slice1)

	fmt.Println(slices.Equal(slice, slice1))

}