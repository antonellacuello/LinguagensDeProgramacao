/*2. Dado um inteiro positivo n, verificar se n é primo*/
#include <stdio.h>

int main() {
    int n, i, primo = 1; // Assume que é primo (1 = verdadeiro)

    printf("Digite um número inteiro positivo: ");
    scanf("%d", &n);

    if (n <= 1) {
        primo = 0; // 0 e 1 não são primos
    } else {
        for (i = 2; i <= n / 2; i++) {
            if (n % i == 0) {
                primo = 0; // Achou um divisor
                break;        // Não precisa continuar
            }
        }
    }

    if (primo) {
        printf("%d é primo.\n", n);
    } else {
        printf("%d não é primo.\n", n);
    }

    return 0;
}
