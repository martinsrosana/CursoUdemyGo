package main

import "fmt"

// ela é executada antes da main e ela pode ter uma em cada pacote
func init() {
	fmt.Println("Executando a função Init")
}

func main() {
	fmt.Println("Executando a função Main")
}
