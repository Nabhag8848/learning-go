package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error){
	if arg == 0 {
		return -1, errors.New("its zero")
	}

	return arg, nil
}

var NotFoundError = errors.New("404 Not Found")
var InternalServerError = errors.New("500 Internal Server Error")

func api(val int) error {
	if val == 500 {
		return InternalServerError
	}

	return fmt.Errorf("not found: %w", NotFoundError)
}

func Errors() {

	fmt.Println(f(1))
	fmt.Println(f(0))
	err := api(500)
	fmt.Println(err)
	if errors.Is(err, InternalServerError) {
		fmt.Println(500)
	}
	fmt.Println(api(404))

}