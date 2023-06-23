package main

import "fmt"

func VerificarParImpar(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}

func main() {
	numero := 7
	fmt.Println("Número antes da verificação:", numero)

	VerificarParImpar(&numero)
	fmt.Println("Número após a verificação:", numero)
}
