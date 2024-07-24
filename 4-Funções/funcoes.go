package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// qdo nao quero usar os dois retornos de uma funcao que tenha dois retornos como o caso acima, usarei o _ no lugar do retorno que nao vou usar
	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)

}
