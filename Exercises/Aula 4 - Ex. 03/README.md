# Aula 4 - exercício 3

Usando função faça um algoritmo que calcula a média de uma lista de números.

```go
package main
import "fmt"

func main() {
    fmt.Println("Lista com os numeros : 1, 2, 3, 4")
    numeros := []float64{1, 2, 3, 4}
    
    media := calcularMedia(numeros)
    fmt.Println("A media = ", media)
    
}

func calcularMedia(numeros []float64) float64{
    total := 0.0
    
    for i := 0; i < len(numeros); i++ {
        total += numeros[i]
    }
    return(total/float64(len(numeros)))
}
```
