package main

import "fmt"

type myType5 int

var x5 myType5

var y5 int

func main() {

	fmt.Printf("x4: %v %T\n", x5, x5)

	x5 = 42
	fmt.Println(x5)

	y5 = int(x5)
	fmt.Printf("y5: %v %T\n", y5, y5)

}
