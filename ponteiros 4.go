package main

import "fmt"

func updateLastDigitsSum(num *int) {

	lastDigit := *num % 10
	secondLastDigit := (*num / 10) % 10

	sum := lastDigit + secondLastDigit

	*num = sum
}

func main() {

	value := 1234
	fmt.Println("Valor inicial:", value)

	updateLastDigitsSum(&value)
	fmt.Println("Novo valor:", value)
}
