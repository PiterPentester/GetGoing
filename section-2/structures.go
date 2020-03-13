package main

import (
	"fmt"
)

//structure encapsulation
type Car struct {
	Name  string
	Age   int
	Model int
}

func (c Car) print() {
	fmt.Println(c)
}

func (c Car) drive() {
	fmt.Println("wroom")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	c := Car{
		Name:  "Jeep",
		Age:   1,
		Model: 1,
	}

	fmt.Println(c)
	c.print()
	c.drive()
	fmt.Println(c.GetName())
}
