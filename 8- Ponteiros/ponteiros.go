package main

import "fmt"

func main() {

	//valores por cópia
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println("Valores por cópia")

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//Ponteiros é uma referencia de memória
	fmt.Println("Ponteiros")

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	fmt.Println(variavel3, *ponteiro) //desreferenciação

	variavel3 = 150

	fmt.Println(variavel3, ponteiro)

	fmt.Println(variavel3, *ponteiro) //desreferenciação

}
