package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Rosana", "Martins", 40, 154}
	fmt.Println(p1)

	e1 := estudante{p1, "Direito", "Uninove"}
	fmt.Println(e1)

}
