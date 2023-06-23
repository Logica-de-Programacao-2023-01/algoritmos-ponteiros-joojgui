package main

import (
	"fmt"
	"math"
)

func updateWithPiAverage(num *float64) {

	currentValue := *num

	average := (currentValue + math.Pi) / 2.0

	*num = average
}

func main() {

	value := 3.14
	fmt.Println("Valor inicial:", value)

	updateWithPiAverage(&value)
	fmt.Println("Novo valor:", value)
}
