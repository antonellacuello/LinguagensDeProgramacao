#include <stdio.h>

int main()
{
    int n, i, fatorial;
    printf("Digite o número que o fatorial deve ser calculado: ");
    scanf("%d", &n);
    i = 2;
    fatorial = 1;
    while (i <= n) {
        fatorial = fatorial * i;
        i++;
    }
    printf("O valor de %d! é %d.", n, fatorial);
    return 0;
}
