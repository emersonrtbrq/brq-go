package main

import "fmt"

func Ex7() {
	// Inicializar três variáveis booleanas
	a := true
	b := false
	c := true

	// Operações lógicas combinadas
	fmt.Println("Valores das variáveis:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	// Operação AND lógico (&&)
	fmt.Println("\nOperações lógicas AND (&&):")
	fmt.Printf("a && b: %v\n", a && b) // false
	fmt.Printf("a && c: %v\n", a && c) // true
	fmt.Printf("b && c: %v\n", b && c) // false

	// Operação OR lógico (||)
	fmt.Println("\nOperações lógicas OR (||):")
	fmt.Printf("a || b: %v\n", a || b) // true
	fmt.Printf("a || c: %v\n", a || c) // true
	fmt.Printf("b || c: %v\n", b || c) // true

	// Operação NOT lógico (!)
	fmt.Println("\nOperações lógicas NOT (!):")
	fmt.Printf("!a: %v\n", !a) // false
	fmt.Printf("!b: %v\n", !b) // true
	fmt.Printf("!c: %v\n", !c) // false
}
