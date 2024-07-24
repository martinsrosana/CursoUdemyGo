package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 2
	fmt.Println(array1)

	array2 := [5]string{"Posicão 1, ", "Posicão 2, ", "Posicão 3, ", "Posicão 4, ", "Posicão 5."}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	fmt.Println("Slice")

	slice := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice)

	//funçao append do slide
	slice = append(slice, 17)
	fmt.Println(slice)

	//reflect.TypeOf devolve o tipo de uma variavel
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	// arrays internos, com slice e função make para alocar um espaço na memoria para uma determinada coisa
	//3 parametros (tipo, tamanho, capacidade máxima )
	fmt.Println("_________")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(len(slice3)) //lenght/tamanho
	fmt.Println(cap(slice3)) //capacidade

}
