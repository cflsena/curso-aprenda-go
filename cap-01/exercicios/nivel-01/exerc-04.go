package main

import "fmt"

type myType4 int

var x4 myType4

func main() {

	fmt.Printf("x4: %v %T\n", x4, x4)

	x4 = 42
	fmt.Println(x4)

}
