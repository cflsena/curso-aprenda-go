package main

import "fmt"

func main() {

	// Contém caracteres que precisam ser interpretados (\n, \t) dentro de uma função
	// como no caso da função print
	interpretedString := "Olá\n, Tudo bom?\tFulano"
	fmt.Println(interpretedString)

	// Será tratada conforme foi declara, sem qualquer tipo de interpretação
	rawString := `"Olá\n, Tudo bom?\tFulano"`
	fmt.Println(rawString)
}
