package main

import "fmt"

// copia do valor, usada qdo vc quer manter o valor original e fazer somente uma cópia
func inverterSinal(numero int) int {
	return numero * -1
}

// referencia do valor, qdo vc não quer manter o valor inicial de uma variavel
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	fmt.Println("Sem mexer no endereço de memória, fazendo uma cópia")
	numero := 20
	fmt.Println(numero)
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	fmt.Println("Mexendo no endereço de memória com PONTEIRO")
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
