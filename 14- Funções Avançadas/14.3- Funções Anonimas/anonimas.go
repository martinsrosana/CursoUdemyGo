package main

import "fmt"

func main() {

	func() {
		fmt.Println("Ola Mundo!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando o Parâmetro")
	// o Stringf vc consegue passar uma string dentro dele e ir concatenando com outros tipos de dados, %s= string, %d = numero
	retorno := func(texto2 string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto2, 10)
	}("Passando parametro")

	fmt.Println(retorno)
}
