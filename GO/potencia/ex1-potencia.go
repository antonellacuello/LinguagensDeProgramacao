/* Dados um inteiro x e um inteiro não-negativo n, calcular 𝑥^n */

package main

import (
	"fmt"
)

func potencia(base, exp int) int {
	if exp == 0 {
		return 1
	} else if exp == 1 {
		return base
	} else {
		return base * potencia(base, exp - 1)
	}
}

func main() {
	var base, exp int 
	fmt.Print("Digite a base e o expoente da potência a ser calculada: ")
	fmt.Scan(&base, &exp)
	resultado := potencia(base, exp)
	fmt.Printf("Resultado da potência de %d^%d: %d\n", base, exp, resultado)
}