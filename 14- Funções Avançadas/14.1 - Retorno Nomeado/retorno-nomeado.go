package main

import "fmt"

// retorno nomeado não precisa colocar o return pq ele já entende que é a soma e a subtraçao
func calculosMatemático(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {

	soma, subtracao := calculosMatemático(10, 20)
	fmt.Println(soma, subtracao)

}
