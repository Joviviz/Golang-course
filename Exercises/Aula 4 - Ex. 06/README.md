# Aula 4 - Exercício 6

Usando função faça um algoritmo que calcula a potência de um número, elevando-o a uma determinada potência.

```go
package main
import "fmt"

func main() {
    var n, potencia float64
    
    fmt.Println("Escreva um numero: ")
    fmt.Scanln(&n)
    
    fmt.Println("Potencia do numero: ")
    fmt.Scanln(&potencia)
    
    fmt.Println("Resultado: ",calcularPotencia(n, potencia))
}

func calcularPotencia(n, potencia float64) float64 {
    total := 1.0
    
    for i:=1.0 ; i <= potencia; i++{
        total *= n
    }
    
    return total
}
```
