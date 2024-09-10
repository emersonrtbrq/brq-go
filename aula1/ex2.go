package main

import "fmt"

func Ex2() {
	// Declarar e inicializar variáveis de diferentes tipos básicos

	// Inteiros
	var i8 int8 = -127
	var i16 int16 = -32767
	var i32 int32 = -2147483647
	var i64 int64 = -9223372036854775807
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var uintptrVar uintptr = 123456789

	// Floats
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793

	// Números Complexos
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 1 + 2i

	// Booleano
	var b bool = true

	// String
	var s string = "Olá, Go!"

	// Rune (equivalente a char)
	var r rune = 'A'

	// Byte (equivalente a uint8)
	var by byte = 'B'

	// Ponteiro
	var x int = 10
	var ptr *int = &x

	// Imprimir as variáveis
	fmt.Println("Informações das variáveis:")
	fmt.Printf("Int8: %d\n", i8)
	fmt.Printf("Int16: %d\n", i16)
	fmt.Printf("Int32: %d\n", i32)
	fmt.Printf("Int64: %d\n", i64)
	fmt.Printf("Uint8: %d\n", u8)
	fmt.Printf("Uint16: %d\n", u16)
	fmt.Printf("Uint32: %d\n", u32)
	fmt.Printf("Uint64: %d\n", u64)
	fmt.Printf("uintptr: %d\n", uintptrVar)

	fmt.Printf("Float32: %f\n", f32)
	fmt.Printf("Float64: %f\n", f64)

	fmt.Printf("Complex64: %v\n", c64)
	fmt.Printf("Complex128: %v\n", c128)

	fmt.Printf("Booleano: %v\n", b)
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Rune: %c\n", r)
	fmt.Printf("Byte: %c\n", by)
	fmt.Printf("Ponteiro: %v (valor apontado: %d)\n", ptr, *ptr)

	// Declarar uma variável sem inicializá-la e imprimir o valor padrão
	var semInicializacao int
	fmt.Println("\nValor zero padrão para variável sem inicialização (int):")
	fmt.Println("Sem Inicialização (int):", semInicializacao)

	// Declarar uma variável sem inicializá-la e imprimir o valor padrão string
	var semInicializacaoString string
	fmt.Println("\nValor zero padrão para variável sem inicialização (string):")
	fmt.Println("Sem Inicialização (string):", semInicializacaoString)
}
