# Quesões Médias

# Q1
```go
package main

import "fmt"

func encontrarNumero(numeros []int) int {
	regIndex := 0

	for i := 0; i < len(numeros)-1; i++ {
		if numeros[i] != numeros[i+1] {
			regIndex = numeros[i]
		}
	}
	return regIndex
}

func main() {

	fmt.Println("Lista com os numeros : 1, 1, 1, 4, 0, 0")
	numeros := []int{1, 1, 1, 4, 0, 0}
	println("O número que não se repete na lista é: ", encontrarNumero(numeros))
}
```
