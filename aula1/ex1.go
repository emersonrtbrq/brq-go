package main

import "fmt"

func Ex1() {
	// Definir variáveis
	nome := "Emerson"
	idade := 25
	cidade := "Santa Cruz do Sul"

	// Usar fmt.Println para imprimir as informações
	fmt.Println("Informações usando fmt.Println:")
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Cidade Natal:", cidade)

	// Usar fmt.Printf para formatar e imprimir as informações
	fmt.Println("\nInformações usando fmt.Printf:")
	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d\n", idade)
	fmt.Printf("Cidade Natal: %s\n", cidade)
}
