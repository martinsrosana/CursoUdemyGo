package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
		fallthrough
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaSemana string
	switch {
	case numero == 1:
		diaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaSemana = "Segunda"
	case numero == 3:
		diaSemana = "terça-feira"
	case numero == 4:
		diaSemana = "Quarta-feira"
	case numero == 5:
		diaSemana = "Quinta-feira"
	case numero == 6:
		diaSemana = "Sexta-feira"
	case numero == 7:
		diaSemana = "Sábado"
	default:
		diaSemana = "Numero inválido"
	}
	return diaSemana
}

func main() {
	fmt.Println("Switch")

	//funçao dia:
	dia := diaDaSemana(1)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)

}
