/* Dados um inteiro x e um inteiro não-negativo n, calcular 𝑥^n */

package potencia

func Calcular(base, exp int) int {
	if exp == 0 {
		return 1
	} else if exp == 1 {
		return base
	} else {
		return base * Calcular(base, exp - 1)
	}
}