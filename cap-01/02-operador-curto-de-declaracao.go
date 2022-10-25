package main

import "fmt"

// Visível no escopo do pacote
var m = "Test"

func main() {

	// Operador curto, declara e atribui ao mesmo tempo. So pode ser utilizado na delcaração de uma variável
	x := 10
	y := "Olá, bom dia!"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x = 20
	fmt.Printf("x: %v, %T\n\n", x, x)

	// Neste caso, o operador servirá como operador de atribuição para "x" (pois ja foi delcarado antes)
	// e atribuição e declaração para "z"
	// Isso só funciona porque estou declarando uma nova variável, ao passo que estou atribuindo um novo valor para "x"
	x, z := 30, 40
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n\n", z, z)

	x, y, w := 50, "olá, boa tarde", 60
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("w: %v, %T\n\n", w, w)

	// Atribuindo novo valor para uma variável de escopo do pacote
	m = "Teste 2"
	fmt.Printf("y: %v, %T\n\n", m, m)

}
