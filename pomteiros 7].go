package main

import "fmt"

type Conta struct {
	saldo   float64
	titular string
}

func adicionarValor(conta *Conta, valor float64) {
	conta.saldo += valor
}

func main() {

	conta := Conta{
		saldo:   100.0,
		titular: "João",
	}

	fmt.Println("Saldo inicial:", conta.saldo)

	valorAdicional := 50.0
	adicionarValor(&conta, valorAdicional)
	fmt.Println("Saldo atualizado:", conta.saldo)
}