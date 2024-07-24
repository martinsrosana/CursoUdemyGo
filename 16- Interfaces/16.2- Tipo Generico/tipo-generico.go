package main

import "fmt"

// a interface generica aceita qqr tipo de dado, string, numero, booleano, float, struct, interface, etc
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Interface tipo generico")

	generica("String")
	generica(1)
	generica(true)

	// o Println recebe uma interface vazia e por isso vc pode passar qtos parametros vc quiser e de qts tipos vc quiser... ex:
	fmt.Println("string", 1, 2, false, true, float64(12345))

	fmt.Println("Nem sempre é indicado udar a interface generica, exemplo do map:")
	//mas vc precisa tomar cuidado com o tipo de interface generica nem sempre é aconselhavel usar:
	//ex no map pq vc passar uma chave [] de qualquer tipo e um valor tbm de qqr tipo usando o interface{} pode virar bagunça
	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapa)
}
