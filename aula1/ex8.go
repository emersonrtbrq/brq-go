package main

import "fmt"

func Ex8() {
	// Declarar uma variável inteira
	var original int = 42

	// Declarar um ponteiro para a variável original
	var ponteiro *int = &original

	// Imprimir o valor da variável original antes da alteração
	fmt.Println("Valor original antes da alteração:", original)

	// Usar o ponteiro para alterar o valor da variável original
	*ponteiro = 100

	// Imprimir o valor da variável original depois da alteração
	fmt.Println("Valor original depois da alteração:", original)
}
