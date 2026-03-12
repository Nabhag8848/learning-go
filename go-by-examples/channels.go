package main

import (
	"fmt"
	"sync"
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

func channelClosed() {
	jobs := make(chan int)
	done := make(chan bool)


	go func() {	
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
                done <- true
                return
			}
		}
	
	}()
	
	for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)

    }
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
    fmt.Println("received more jobs:", ok)
}

func rangeOverChannel(){
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for i := range queue {
		fmt.Println(i)
	}
}

func workerr(i int){
	fmt.Println("worker sleep - ", i)
	time.Sleep(1 * time.Second)
	fmt.Println("worker executed and return - ", i)
}

func WaitGroup() {
	var wg sync.WaitGroup

	for i:= range 5 {
		wg.Go(func() {
			workerr(i)
		})
	}

	wg.Wait()
}

func Channels() {
	sendValuesOverChannel()
	channelBuffering()
	channelSynchronization()
	channelDirections()
	channelValueSelect()
	channelClosed()
	rangeOverChannel()
	WaitGroup()
}