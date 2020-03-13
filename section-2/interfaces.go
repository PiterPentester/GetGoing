/*
package main

import (
	"fmt"
)

type Car interface {
	Drive()
	Stop()
}

type Jeep struct {
	JeepModel string
}

type Lada struct {
	LadaModel string
}

func (j *Jeep) Drive() {
	fmt.Println("Jeep: wroom wroom")
	fmt.Println(j.JeepModel)
}

func (l *Lada) Drive() {
	fmt.Println("Lada: wroom wroom")
	fmt.Println(l.LadaModel)
}

func (j *Jeep) Stop() {
	fmt.Println("Jeep: stop")
}

func (l *Lada) Stop() {
	fmt.Println("Lada: stop")
}

// ENFORCE interface
func NewModel(arg string) Car {
	return &Jeep{arg}
}

func main() {
	l := Lada{"Niva"}

  //test NewModel(arg string)
  j := NewModel("ZED")
	l.Drive()
	l.Stop()
	j.Drive()
	j.Stop()

}
*/

package main

import (
	"fmt"
)

func Anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	Anything("WOW")
	mymap := make(map[string]interface{})
	mymap["name"] = "devoops"
	mymap["age"] = 21
	fmt.Println(mymap)

	//test interface
	mymap2 := make(map[interface{}]interface{})
	mymap2[34] = "devoops"
	mymap2["age"] = 21
	fmt.Println(mymap2)
}
