package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Go!")
	}
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Go super heavy!")
	}
}

func main() {
	go heavy()
	go superHeavy()
	fmt.Println("Finish!")
	time.Sleep(time.Second * 5)
}
