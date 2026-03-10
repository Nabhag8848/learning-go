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

type argError struct {
    arg     int
    message string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f1(arg int) (int, error) {
    if arg == 42 {

        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
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

	_, err1 := f1(42)
    if ae, ok := errors.AsType[*argError](err1); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
}



    