# Aula 3 - Exercício 5

Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.

```go
package main

import "fmt"

func main() {
	var n int
	println("Digite um numero inteiro positivo: ")
	fmt.Scan(&n)

	algo(n)
}

func algo(n int) {
	for i := 1; i < n+1; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
```
