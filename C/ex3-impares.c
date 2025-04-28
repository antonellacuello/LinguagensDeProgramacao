/* 3. Dado um número inteiro positivo n, imprimir os n primeiros naturais ímpares.*/

#include <stdio.h>

int main()
{
    int n, i;
    printf("Digite um número inteiro positivo: ");
    scanf("%d", &n);
    
    printf("Os primeiro naturais ímpares são: ");
    for (i = 1; i <= n; i++) {
        if (i % 2 != 0) {
            printf("\n%d", i);
        }
    }
    return 0;
}
