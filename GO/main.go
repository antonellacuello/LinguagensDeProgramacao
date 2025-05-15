package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"go-linguagens/potencia"
	"go-linguagens/primo"
    "go-linguagens/impares"
    "go-linguagens/conversaotemp"
	"go-linguagens/ocorrencias"
)


func main() {
	// EX1 - Potência
	fmt.Println("EXERCÍCIO 1")
	var base, exp int
	fmt.Print("Digite a base e o expoente: ")
	fmt.Scan(&base, &exp)
	resultado := potencia.Calcular(base, exp)
	fmt.Printf("Resultado da potência de %d^%d: %d\n", base, exp, resultado)

	// EX2 - Primo
	fmt.Println("----------------------------------------")
	fmt.Println("EXERCÍCIO 2")
	var n int
	fmt.Print("Digite o número que deseja saber se é primo: ")
	fmt.Scan(&n)
	teste := primo.Testar(n)
	if teste {
		fmt.Println(n, "é primo.")
	} else {
		fmt.Println(n, "não é primo.")
	}

	// EX3 - Ímpares
	fmt.Println("----------------------------------------")
	fmt.Println("EXERCÍCIO 3")
	var n1 int
	fmt.Print("Digite um número para ver os ímpares: ")
	fmt.Scan(&n1)
	impares.Imprimir(n1)

	// EX4 - Conversão DE Temperatura
	fmt.Println("----------------------------------------")
	fmt.Println("EXERCÍCIO 4")
	var tempF int
	fmt.Print("Insira a temperatura em Fahrenheit: ")
	fmt.Scan(&tempF)
	conversao := conversaotemp.Converter(tempF)
	fmt.Println(tempF, "°F em °C é", conversao)

	// EX5 - Ocorrência de Palavra
	fmt.Println("----------------------------------------")
	fmt.Println("EXERCÍCIO 5")
	var frase, palavra string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Print("Digite a palavra: ")
	palavra, _ = leitor.ReadString('\n')
	palavra = strings.TrimSpace(palavra)
	palavra = strings.ToUpper(palavra)

	fmt.Print("Digite a frase: ")
	frase, _ = leitor.ReadString('\n')
	frase = strings.TrimSpace(frase)
	frase = strings.ToUpper(frase)

	ocorrencias.Ocorrencias(frase, palavra)
}