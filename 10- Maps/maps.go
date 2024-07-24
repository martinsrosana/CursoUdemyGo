package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// dentro do[tipoda chave] e fora é o tipo do valor da chave
	// os dois campos tem que ser todos do mesmo tipo, chave e valor. Nao consigo usar uma chave string e a outra int
	usuario := map[string]string{
		"nome":      "Rosana",
		"sobrenome": "Martins",
	}

	fmt.Println(usuario)

	// para acessar não consigo fazendo usuario.nome preciso fazer igual esse println abaixo:

	fmt.Println(usuario["nome"])

	//map aninhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	// nome e cuso são chaves, caso queira deletar uma chave uso a funçao delete do go

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	// criar uma nova chave com só um valor tbm é possivel, no mesmo map
	usuario2["signo"] = map[string]string{
		"nome": "Peixes",
	}

	fmt.Println(usuario2)
}
