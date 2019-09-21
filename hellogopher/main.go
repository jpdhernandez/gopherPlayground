package main

import (
	"fmt"

	"github.com/jpdhernandez/stringutils"
)

const state = "Toronto"

var name string

func main() {

	name = "julian"
	from := `Philippines`
	var n int

	fmt.Printf("Hello, fellow %s Gophers!\n", state)
	fmt.Printf("My name is %s and I'm from %s.\n", stringutils.Upper(name), from)
	fmt.Printf("By the time %d o'clock comes around, we'll know how to code in Go!\\n", n)
	fmt.Println("Let's get started!")
}
