package main

import "fmt"

func main() {

	// Print -> imprime o valor no console -> Standard out - format verbs %v %T
	x := 10
	fmt.Print("Teste\n")
	fmt.Println("Teste 2")
	fmt.Printf("Teste %v\n", 3)
	fmt.Printf("x = %v %T\n\n", x, x)

	// Sprint (String print) -> retorna uma string de acordo com os parametros passados
	result := fmt.Sprintln(1, 2)
	fmt.Println(result)
	result2 := fmt.Sprintln("1", "2")
	fmt.Println(result2)

	// Fprint (File print) -> necessita receber como parâmetro um interface de escrita (io.Writer)

}
