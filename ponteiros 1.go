package main

import "fmt"

func SomarNumerosNaturais(ptr *int, n int) {
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	*ptr = soma
}

func main() {
	numero := 0
	n := 5
	SomarNumerosNaturais(&numero, n)
	fmt.Println("Soma dos", n, "primeiros números naturais:", numero)
}
