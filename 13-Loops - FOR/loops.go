package main

import (
	"fmt"
	//"time"
)

func main() {
	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// range para interar com strings, maps...

	fmt.Println("_____RANGE______")

	nomes := [3]string{"Joana", "Marcia", "Camila"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// se não quiser exibir o indice do array preciso usar o _ assim:

	fmt.Println("_________")
	fmt.Println("Sem o indice do array")
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	// obs o struct não aceita range, somente o map e os demais exemplos acima aceitam o range
	usuario := map[string]string{
		"nome":      "Rosana",
		"Sobrenome": "Martins",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// loop infinito

	// for {
	// 	fmt.Println("For infinito")
	// 	time.Sleep(time.Second)
	// }
}
