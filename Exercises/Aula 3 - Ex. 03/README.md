# Aula 3 - Exercício 3

Faça um algoritmo que imprima a tabuada de multiplicação de 1 a 10 para um número fornecido pelo usuário.

```go
package main
import "fmt"

func main() {
	var n float64
	fmt.Print("Escolha um numero: \n")
	fmt.Scan(&n)

	tabuada(n)
}

func tabuada(n float64) {
	var i float64
	i = 1
	for i < 11 {

		fmt.Printf("%.2f x %.2f = %.2f \n", n, i, n*i)
		i++
	}
}
```
