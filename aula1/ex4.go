package main

import "fmt"

func Ex4() {
	// Instanciar dois números inteiros
	var num1 int = 15
	var num2 int = 4

	// Realizar as operações
	soma := num1 + num2
	subtracao := num1 - num2
	produto := num1 * num2
	quociente := num1 / num2
	modulo := num1 % num2

	// Imprimir os resultados
	fmt.Println("Operações com os números", num1, "e", num2)
	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Produto:", produto)
	fmt.Println("Quociente:", quociente)
	fmt.Println("Resto da divisão (módulo):", modulo)
}
