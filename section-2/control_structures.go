package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	//if else ,for , switch, break continue
	f := true
	flag := &f
	if flag == nil {
		fmt.Println("Empty flag!!!")
	} else if !*flag {
		fmt.Println("false")
	} else {
		fmt.Println("True")
	}

	// FOR
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// in arr range by value
	arr := []string{"Wow", "cool", "I am a Progger))"}
	for _, v := range arr {
		fmt.Println(v)
	}

	//switch, continue, break
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	flag1 := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag1 = false
			//break
		} else if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println(flag1)

	// switch
	d := "Fri"
	switch d {
	case "Fri":
		fmt.Println("FRIDAY")
	case "Mon":
		fmt.Println("MONDAY")
	case "Tue":
		fmt.Println("TUESDAY")
	default:
		fmt.Println("OK")
	}

}
