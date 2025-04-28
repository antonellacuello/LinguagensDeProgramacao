/* 4. Escreva um programa que converta temperaturas de Fahrenheit para Celsus, seguindo
a equação 𝐶 = (𝐹−32) * 5 / 9    */

#include <stdio.h>

int main()
{
    int f, c;
    printf("Digite a temperatura em Fahrenheit: ");
    scanf("%d", &f);

    c = ((f - 32) * 5) / 9;

    printf("%dºF é %dºC.", f, c);
    return 0;
}