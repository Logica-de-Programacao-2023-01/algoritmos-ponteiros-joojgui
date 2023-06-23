package main

import "fmt"

func preencherPrimos(slice *[]int, tamanho int) {
	if tamanho <= 0 {
		return
	}

	primosEncontrados := 0
	numero := 2

	for primosEncontrados < tamanho {
		if ehPrimo(numero) {
			*slice = append(*slice, numero)
			primosEncontrados++
		}
		numero++
	}
}

func ehPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func main() {

	var primos []int
	tamanho := 10

	preencherPrimos(&primos, tamanho)
	fmt.Println("Números primos:", primos)
}
