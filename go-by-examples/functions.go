package main

import "fmt"

func Functions() {
	a, b, c := multipleReturnValues()
	fmt.Println(a,b,c)

	anyNumberOfArguments(1,2,3,4,5,)
	anyNumberOfArguments()
	anyNumberOfArguments(1, 2)
	anyNumberOfArguments([]int{1,2,4,5}...)

	incrementer := counter()
	fmt.Println(incrementer())
	fmt.Println(incrementer())
	fmt.Println(incrementer())
	fmt.Println(incrementer())
	fmt.Println(incrementer())
	fmt.Println(incrementer())

	anotherIncrementer := counter()
	fmt.Println(anotherIncrementer())
	fmt.Println(anotherIncrementer())
	fmt.Println(anotherIncrementer())

	recursiveClosure()

}

func multipleReturnValues() (int, int, int) {
	return 1 ,2, 3
}

func anyNumberOfArguments(nums ...int) {
	fmt.Println(nums)
}

func counter() func () int {
	var i int = 0

	return func () int {
		i++
		return i
	}
}

func recursiveClosure() {
	var fact func (n int) int

	fact = func (n int) int {
		if n < 2 {
			return n
		}

		return n * fact(n - 1)
	}

	fmt.Println(fact(5))
}