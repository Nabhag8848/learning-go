package main

import (
	"fmt"
	"time"
)

func Timeouts() {
	c := make(chan string, 1)

	go func () {
		time.Sleep(2 * time.Second)
		c <- "done"
	}()

	select {
	case msg:= <-c: 
		fmt.Print(msg)
	case <-time.After(2 * time.Second): 
		fmt.Println("timeout 2")
	}
}