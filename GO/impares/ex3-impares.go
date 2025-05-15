/* Dado um número inteiro positivo n, imprimir os n primeiros naturais ímpares. */
package impares

import (
	"fmt"
)

func Imprimir(n1 int) {
	fmt.Print("Os primeiros naturais ímpares são: ")
	for i := 1; i <= n1; i++ {
		if i % 2 != 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
