package main

import "fmt"

type milkshake int

var c milkshake

func main() {

	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", c)

	// Em Go ha somente conversao de tipos, nao se diz casting
	c = 20
	x = int(c)
	fmt.Println(x)

}
