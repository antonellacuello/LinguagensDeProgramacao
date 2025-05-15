package main

import (
	"fmt"
	"GO/impares"  
	"GO/potencia"
)

func main() {
	// EX1 - Potência
	var base, exp int
	fmt.Print("Digite a base e o expoente: ")
	fmt.Scan(&base, &exp)
	resultado := potencia.Calcular(base, exp)
	fmt.Printf("Resultado da potência de %d^%d: %d\n", base, exp, resultado)

	// EX2 - Ímpares
	var n int
	fmt.Print("Digite um número para ver os ímpares: ")
	fmt.Scan(&n)
	impares.Imprimir(n)