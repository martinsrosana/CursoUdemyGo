package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else if numero < -10 {
		fmt.Println("Numero é menor que - 10")
	} else {
		fmt.Println("Está entre 0 e -10")
	}
}
