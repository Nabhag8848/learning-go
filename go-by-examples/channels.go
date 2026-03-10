package main

import (
	"fmt"
	"time"
)

func sendValuesOverChannel() {
	messages:= make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

func channelBuffering() {
	c := make(chan string, 2)

	c <- "buffered"
	c <- "channel"

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func worker(done chan bool) {
	fmt.Println("working..")
	time.Sleep(time.Second)
	fmt.Println("done..")

	done <- true
}

func channelSynchronization() {
	channel := make(chan bool)
	go worker(channel)

	<-channel
}

func pings(pings chan<- string){
	pings <- "ping"
}

func pongs(pings <-chan string, pongs chan<- string){
	msg := <- pings
	fmt.Println(msg)
	pongs <- "pong"
}
func channelDirections() {
	ping := make(chan string , 1)
	pong := make(chan string , 1)

	pings(ping)
	pongs(ping, pong)
	fmt.Println(<-pong)
}

func channelValueSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func () {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func () {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
		}
	}
}

func Channels() {
	sendValuesOverChannel()
	channelBuffering()
	channelSynchronization()
	channelDirections()
	channelValueSelect()
}