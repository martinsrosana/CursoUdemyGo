package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// esse metodo ainda não salva em BD vms ver lá na frente
func (user usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s \n", user.nome)
}

// metodo pra saber se o usuário é maior de idade, aqui sem ponteiro faz somente uma cópia
func (user usuario) maiorDeIdade() bool {
	return user.idade >= 18
}

// metodo que altera a idade, usando ponteiro para as alteraçoes persistirem em todos os lugares que forem usadas
func (user *usuario) fazerAniversario() {
	user.idade++
}

func main() {
	fmt.Println("Método")

	usuario1 := usuario{"Rosana", 40}
	usuario1.salvar()
	maiorIdade1 := usuario1.maiorDeIdade()
	fmt.Println(maiorIdade1)
	fmt.Println(usuario1)

	usuario2 := usuario{"Artemis", 17}
	usuario2.salvar()
	maiorIdade2 := usuario2.maiorDeIdade()
	fmt.Println(maiorIdade2)
	fmt.Println(usuario2)
	fmt.Printf("Metodo Fazendo aniversário, idade Usuário %s \n", usuario2.nome)
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
