package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1 ")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {

	defer fmt.Println("Media Calculada. Resultado será retornado!")
	fmt.Println("Entrando a funça para verificar se o aluno esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false

}

func main() {
	fmt.Println("Defer")
	defer funcao1()
	// defer em ingles é adiar entao ele adia até o ultimo momento, ele é util para manipular BD e ou para reornar um print de uma funcao...
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))

}
