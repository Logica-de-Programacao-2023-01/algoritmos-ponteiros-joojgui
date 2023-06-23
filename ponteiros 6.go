package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
}

func atualizarTitulo(livro *Livro) {
	if livro.autor == "Anônimo" {
		livro.titulo = "Desconhecido"
	}
}

func main() {

	livro := Livro{
		titulo: "Livro X",
		autor:  "Anônimo",
	}

	fmt.Println("Título inicial:", livro.titulo)

	atualizarTitulo(&livro)
	fmt.Println("Novo título:", livro.titulo)
}