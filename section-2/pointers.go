package main

import (
	"fmt"
)

func swap(p1, p2 *int) {
	var tmp int
	tmp = *p2
	*p2 = *p1
	*p1 = tmp
}

func main() {
	p1 := 2
	p2 := 4
	pointer := &p1

	fmt.Println(pointer)
	//get value of a pointer
	fmt.Println(*pointer)

	//swap values
	fmt.Println("p1 =", p1, "p2 =", p2)
	swap(&p1, &p2)
	fmt.Println("Swap values")
	fmt.Println("p1 =", p1, "p2 =", p2)

}
