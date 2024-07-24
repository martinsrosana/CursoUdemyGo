package main

import "fmt"

// nao tem classe em go mas tem o structs, ele é uma colecao de campos

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua José de Antonio", 6}

	usuario2 := usuario{"Rosana", 40, enderecoExemplo}
	fmt.Println(usuario2)

	// se quiser passar só um atributo do structs faço assim:

	fmt.Println("Passando um atributo por vez do structs")

	usuario3 := usuario{nome: "Artemis"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade: 17}
	fmt.Println(usuario4)

}
