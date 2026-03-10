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

func Channels() {
	sendValuesOverChannel()
	channelBuffering()
	channelSynchronization()
}