package main

import (
	"fmt"
	"time"
)

func Timers() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func () {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()

	// time.Sleep(2 * time.Second) // fired
	// time.Sleep(time.Second) // contradicts so nothing happens for timer2

	stop := timer2.Stop()

	if stop {
		fmt.Println("timer 2 stopped")
	}

}

func Tickers() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
				case <-done:
					return
				case t:= <-ticker.C: 
					fmt.Println("ticker at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("stopped the ticker")
}