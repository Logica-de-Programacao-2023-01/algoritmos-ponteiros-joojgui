package main

import "fmt"

func ReverterString(ptr *string) {
	runes := []rune(*ptr)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*ptr = string(runes)
}

func main() {
	texto := "Hello, world!"
	fmt.Println("Texto antes de reverter:", texto)

	ReverterString(&texto)
	fmt.Println("Texto após reverter:", texto)
}
