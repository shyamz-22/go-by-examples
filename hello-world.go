package main

import (
	"fmt"
	"github.com/go-by-examples/maths"
	"github.com/go-by-examples/arithmetic"
)

// BODMAS
func main() {
	fmt.Println("hello world")
	var number12 = 12

	fmt.Println(maths.AlgebraAPlusBFormula(number12, 33))
	fmt.Println(maths.AlgebraAMinusBFormula(number12, 33))
	fmt.Println(arithmetic.Compute(number12, 33))

	number12 = 13

	//fmt.Println(Algebraaplusbplusc(1, 2, 3))

	printMe("Sai")
	printMe("Karthik")
	printMe("Shyam")

	fmt.Println("########################")

	printAll("Sai", "Shyam", "Karthik", "Uma", "Umamaheswaran", "Pushpa")
}

func printAll(names ...string) {
	for index, name := range names {
		fmt.Println(fmt.Sprintf("At index %d is the name %s", index, name))
		printMe(name)
	}
}

func printMe(name string) {
	fmt.Println(name + " is learning go lang")
	fmt.Println(name + " is watching this is us")
	fmt.Println(name + " is sharing his screen")
}
