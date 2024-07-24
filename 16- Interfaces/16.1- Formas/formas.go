package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type quadrado struct {
	lado float64
}

// como cria uma interface: só passamos as assinaturas dos método, não implemento o método dentro dele
// em GO o implements de uma interface é implícito, não preciso fazer igual em java e usar o implements...

type forma interface {
	area() float64
}

//funçao que vai receber uma estrutura abstrata da interface:

func escreverArea(form forma) {
	fmt.Printf("A area da forma é %0.2f \n", form.area())
}

// preciso criar um método para cada forma do mesmo jeito que esta na assinatura da interface nesse caso, area():

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio) //ou return math.Pi * math.Pow (c.raio, 2)
}

func (q quadrado) area() float64 {
	return math.Pow(q.lado, 2)
}

func main() {
	fmt.Println("Interface")

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

	q := quadrado{10}
	escreverArea(q)
}
