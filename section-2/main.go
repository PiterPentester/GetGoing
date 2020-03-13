package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		a = 1
		b = 3
	)
	fmt.Println(a * b)

	// different data types
	var c int32
	var d int64

	c = 12
	d = 45

	// fmt.Println(c * d) < -- Error, because we have different data types
	fmt.Println(int64(c) * d)

	// short declararion
	x1 := 4
	x2 := 5
	fmt.Println(x1 + x2)

	// STRINGS:
	var s string
	s = "interesting things"
	fmt.Println(s)

	s2 := "things"
	fmt.Println(s2)

	fmt.Println(strings.Contains(s, s2))

	todo()

}
