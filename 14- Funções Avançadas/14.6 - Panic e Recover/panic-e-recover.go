package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}
func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	// panic nao é a melhor forma de tratar um erro e o recover é pra conseguir recuperar a execução qdo se tem um panic
	panic("A Media é exatamente 6!")

}

func main() {
	fmt.Println("Panic e Recover")
	fmt.Println(alunoEstaAprovado(8, 6))
	fmt.Println("Pós execução")

}
