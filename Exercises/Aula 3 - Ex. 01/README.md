# Aula 3 - Exercício 1

Faça um algoritmo que imprima os múltiplos de 3 de 0 a 30.

```go
package main

import "fmt"

func main() {
	for i := 0; i < 31; i++ {
		fmt.Printf("Multiplo de 3 x %d = %d \n", i, 3*i)
	}
}
```
