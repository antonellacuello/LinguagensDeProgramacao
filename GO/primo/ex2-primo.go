/* Dado um inteiro positivo n, verificar se n é primo. */
package primo

import (

)

func Testar(n int) bool {
	primo := true
	if n <= 1 {
		primo = false
	} else {
		for i := 2; i <= n / 2; i++ {
			if n % i == 0 {
				primo = false
				break
			}
		}
	}
	return primo

}