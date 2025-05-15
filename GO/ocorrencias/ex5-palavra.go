/* Dados dois strings (um contendo uma frase e outro contendo uma palavra), determine
o número de vezes que a palavra ocorre na frase. */
package ocorrencias

import (
	"fmt"
	"strings"
)

func Ocorrencias(frase, palavra string) {
	ocorrencias := strings.Count(frase, palavra)

	fmt.Printf("Temos que a palavra %s ocorre %d vez(es) na frase.\n", palavra, ocorrencias)
}