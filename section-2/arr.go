package main

import (
	"fmt"
)

func todo() {
	var arr []int
	arr = []int{1, 2, 3, 4}
	arr2 := []string{"Hey,", "my", "name"}
	//you can append many elelms
	arr2 = append(arr2, "is Bohdan!")
	fmt.Println(arr)
	fmt.Println(arr2)
}

func main() {
	todo()
}
