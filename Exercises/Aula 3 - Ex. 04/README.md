# Aula 3 - Exercício 4

Faça um algoritmo que imprima os números de 1 a 100, substituindo os múltiplos de 3 pela palavra "Fizz" e os múltiplos de 5 pela palavra "Buzz". Para os números que são múltiplos de ambos, utilize a palavra "FizzBuzz".

```go
package main
import "fmt"

func main() {
	fizzbuzz()
}

func fizzbuzz() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz \n")
		} else if i%3 == 0 {
			fmt.Print("Fizz \n")
		} else if i%5 == 0 {
			fmt.Print("Buzz \n")
		} else {
			fmt.Printf("%d \n", i)
		}
	}
}
```
