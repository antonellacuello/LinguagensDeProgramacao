/* Dado um número inteiro positivo n, imprimir os n primeiros naturais ímpares. */
package impares

import (
	"fmt"
)

func Imprimir() {
	var n int
	fmt.Print("Digite um número inteiro e positivo: ")
	fmt.Scanf("%d", &n)
	fmt.Print("Os primeiros naturais ímpares são: ")
	for i := 1; i <= n; i++ {
		if i % 2 != 0 {
			fmt.Print(i)
		}
	}
}