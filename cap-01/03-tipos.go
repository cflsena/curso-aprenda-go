package main

import "fmt"

var z int

func main() {

	// primitivos

	a := 10    // int
	b := "Olá" // string
	c := true  // bool

	fmt.Print(a, b, c, z)
	// tipos compostos -> slice, array, struct, map

	// Valor Zero (valor inicial, antes de ser inicializada)
	// int    -> 0
	// floats -> 0.0
	// bool   -> false
	// string -> ""
	// pointers, functions, interfaces, slices, channels, maps -> nil
}
