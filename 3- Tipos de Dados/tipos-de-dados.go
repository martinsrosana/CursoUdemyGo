package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int12, int64, são os tipos ionteiros do go
	var numero int64 = 10000000000000
	fmt.Println(numero)

	var numero1 uint32 = 10000
	fmt.Println(numero1)

	// alias
	// INT32 = RUNE
	var numero2 rune = 1234
	fmt.Println(numero2)

	// BYTE = UINT8, o uint não aceita numero com sinais
	var numero3 byte = 123
	fmt.Println(numero3)

	//float é sempre .
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	// float por iferencia
	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str1 := "texto1"
	fmt.Println(str1)

	// em go não existe char, o char sempre vai referenciar um numero da tabela asc desse caracter
	char := '%'
	fmt.Println(char)

	// vai retornar vazio
	var texto string
	fmt.Println(texto)

	// vai retornar zero
	var texto1 int16
	fmt.Println(texto1)

	texto2 := 5
	fmt.Println(texto2)

	// bolleano O VALOR SEM INICIALIZAR É FALSE
	var booleano1 bool
	fmt.Println(booleano1)

	var booleano2 bool = true
	fmt.Println(booleano2)

	var erro error = errors.New("Erro interno!")
	fmt.Println(erro)

}
