/* 5. Dados dois strings (um contendo uma frase e outro 
contendo uma palavra), determine o número de vezes que a 
palavra ocorre na frase. */

#include <stdio.h>
#include <string.h>
int main ()
{
    char palavra[20], frase[200];
    int i, j, ocorrencias = 0;
    int tam_frase, tam_palavra;
    int encontrado;
    printf("Digite a frase: ");

    // Remover '\n' da frase, se existir
    fgets(frase, sizeof(frase), stdin); // fgets é uma função para ler uma linha inteira (com espaços) até o Enter (\n).
    printf("Digite a palavra: ");
    scanf("%s", palavra);

    frase[strcspn(frase, "\n")] = '\0';  //strcspn - calcula o número de caracteres consecutivos em s1 até encontrar qualquer caractere que exista em s2calcula o número de caracteres consecutivos em s1 até encontrar qualquer caractere que exista em s2
    tam_frase = strlen(frase);
    tam_palavra = strlen(palavra);

    for (i = 0; i <= tam_frase - tam_palavra; i++) {
        encontrado = 1;

        // verifica se palavra bate a partir da posição i
        for (j = 0; j < tam_palavra; j++) {
            if (frase[i + j] != palavra[j]) {
                encontrado = 0;
                break;
            }
        }

        if (encontrado) {
            ocorrencias++;
        }
    }

    printf("A palavra '%s' ocorre %d vezes na frase.\n", palavra, ocorrencias);

    return 0;
}