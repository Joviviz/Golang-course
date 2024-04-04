# Aula 4 - Exercício 4

Usando função faça um algoritmo que calcula o fatorial de um número, ou seja, o produto de todos os números inteiros de 1 até aquele número.

```go
package main
import "fmt"

func main() {
    n := 3
    fmt.Println("Fatorial de: ", n)
    fmt.Println("O resultado = ", calcularFatorial(n))
    
}

func calcularFatorial(n int) int{
    resultado := 1
    i := 2
    
    for i <= n{
        resultado *= i
        i++
    }
    return resultado
}
```
