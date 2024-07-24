package main

import "fmt"

// Sequencia de fibronacci
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")
	// 1 1 2 3 5 8 13 21 34 55

	posicao := uint(15)
	fmt.Println(fibonacci(posicao))

	fmt.Println("Usando um for para printar a sequencia toda")

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
