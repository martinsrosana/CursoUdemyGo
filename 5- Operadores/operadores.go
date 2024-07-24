package main

import "fmt"

func main() {
	//Aritmeticos

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restosDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restosDaDivisao)

	// qdo for fazer qqr calculo as variaveis precdisam ser do mesmo tipo
	//nao consigo somar int16 com int32

	var numero1 int16 = 10
	var numero2 int16 = 15
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//atribuiçao = ou inferencia de tipo :=

	var variavel1 string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)

	// operadores relacionais, sempre retornam um booleano ou seja true ou false
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// operadores logicos
	// and (&&)
	// ou (||)
	// negaçao (!)

	fmt.Println("___________________________")

	verdadeiro, falso := true, false

	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// operadores unarios
	numero := 10
	numero++
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20
	fmt.Println(numero)

	numero *= 3 // numero = numero * 3
	fmt.Println(numero)

	numero /= 10 // numero = numero / 10
	fmt.Println(numero)

	numero %= 3 // numero = numero / 3 , resto da divisao
	fmt.Println(numero)

	// Operador ternário nao existe em go
	// ex: texto := numero > 5 ? "Maior que 5": "menor que 5"
	// em go nao ;e possivel retornar false ou true para uma mesma variavel, por isso que nao funciona o operador ternário entao a gente usa i if e else:

	fmt.Println("________________")

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
