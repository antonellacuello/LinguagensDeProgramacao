/* Escreva um programa que converta temperaturas de Fahrenheit para Celsius, seguindo a equação 𝐶 = (𝐹−32)×5 / 9 */
package conversaotemp

import (

)

func Converter(tempF int) int {
	tempC := ((tempF - 32) * 5) / 9
	return tempC
}