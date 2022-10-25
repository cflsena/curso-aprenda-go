package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T", b)

	// Neste caso x != b, o fato do tipo hotdog possuir um tipo subjacente int
	// não se aplica regra de transitividade > b == hotdog e hotdog == int entao b == int (isso não possível)
}
