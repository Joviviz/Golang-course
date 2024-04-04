# Aula 4 - Exercício 5

Usando função faça um algoritmo que procura por um determinado elemento em uma lista, verificando cada elemento da lista até encontrá-lo.

```go
package main
import "fmt"

func main() {

    fmt.Println("Lista com os numeros : 1, 2, 3, 4, -1, -2, -3, -4")
    numeros := []float64{1, 2, 3, 4, -1, -2, -3, -4}
    
    encontrarNumero(numeros)
}

func encontrarNumero(numeros []float64) float64{
    var x float64
    
    fmt.Println("Escolha um numero na lista: ")
    fmt.Scanln(&x)
    
    for i := 0; i < len(numeros); i++ {
        if x == numeros[i]{
            fmt.Printf("O numero %.2f esta na lista", numeros[i])
            return x
            break
        } 
    }
    
    fmt.Printf("O numero %.2f nao esta na lista", x)
    return 0
}
```
