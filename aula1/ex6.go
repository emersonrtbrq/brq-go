package main

import "fmt"

func Ex6() {
	// Inicializar dois números inteiros
	num1 := 15
	num2 := 10

	// Comparações usando operadores relacionais
	fmt.Println("Comparações entre", num1, "e", num2)

	fmt.Printf("%d == %d: %v\n", num1, num2, num1 == num2)
	fmt.Printf("%d != %d: %v\n", num1, num2, num1 != num2)
	fmt.Printf("%d < %d: %v\n", num1, num2, num1 < num2)
	fmt.Printf("%d > %d: %v\n", num1, num2, num1 > num2)
	fmt.Printf("%d <= %d: %v\n", num1, num2, num1 <= num2)
	fmt.Printf("%d >= %d: %v\n", num1, num2, num1 >= num2)
}
