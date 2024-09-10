package main

import "fmt"

func Ex3() {
	// Inicializar um número inteiro
	var inteiro int = 42

	// Converter o número inteiro para ponto flutuante
	var flutuante float64 = float64(inteiro)

	// Converter o número inteiro para string
	var str string = fmt.Sprintf("%d", inteiro)

	// Imprimir o valor original e as conversões
	fmt.Println("Número inteiro:")
	fmt.Println("Inteiro:", inteiro)

	fmt.Println("\nConversões:")
	fmt.Printf("Para ponto flutuante: %f", flutuante)
	fmt.Printf("\nPara string: %s\n", str)
}
