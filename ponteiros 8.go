package main

import "fmt"

func modificarValor(ponteiro *int) {

	if ponteiro == nil {
		fmt.Println("Erro: ponteiro nulo")
		return
	}

	*ponteiro = 42

	*ponteiro = 0
}

func main() {

	var valor int
	ponteiro := &valor

	fmt.Println("Valor inicial:", valor)

	modificarValor(ponteiro)
	fmt.Println("Valor modificado:", valor)
}
