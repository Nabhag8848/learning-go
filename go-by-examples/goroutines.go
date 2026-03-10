package main

import (
	"fmt"
	"time"
)


func function(from string) {
	for i := range 3 {
		fmt.Println(from , " : ", i)
	}
}

func GoRoutines() {
	function("direct")


	go function("goroutines")

	go func(msg string) {
		go function(msg)
	}("anon")

	time.Sleep(time.Second)
	fmt.Println("done")
}
