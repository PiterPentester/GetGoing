package main

import (
	"fmt"
	"time"
)

// Pinger func prints "ping" and wait for a "pong"
func pinger(pinger <-chan int, ponger chan int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

// Ponger func prints "pong" and wait for a "ping"
func ponger(pinger chan int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	// Main goroutine starts the ping/pong by sending 1 into the ping chan
	ping <- 1

	for {
		// Block main thread until an interrupt
		time.Sleep(time.Second)
	}

}
