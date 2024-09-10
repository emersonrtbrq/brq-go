package main

import "fmt"

func Ex5() {
	// Inicializar a variável x
	x := 10

	// Imprimir valor inicial
	fmt.Println("Valor inicial de x:", x)

	// Usar operadores de atribuição e exibir o resultado após cada operação
	x += 5
	fmt.Println("Após x += 5:", x)

	x -= 3
	fmt.Println("Após x -= 3:", x)

	x *= 2
	fmt.Println("Após x *= 2:", x)

	x /= 4
	fmt.Println("Após x /= 4:", x)

	x %= 3
	fmt.Println("Após x %= 3:", x)
}
