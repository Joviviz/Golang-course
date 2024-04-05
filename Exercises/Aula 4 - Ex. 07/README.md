# Aula 4 - Exercício 7

Usando função faça um algoritmo que verifica se um número é primo, ou seja, se é divisível apenas por 1 e por ele mesmo.

```go

package main
import "fmt"

func main() {
    var n int
    
    fmt.Println("Escreva um numero: ")
    fmt.Scanln(&n)

    conferirPrimo(n)
    
    switch conferirPrimo(n){
        case true:
            fmt.Printf(" %d = Primo!", n)
        case false:
            fmt.Printf(" %d = Nao primo!", n)
        default:
            fmt.Printf(" %d = Erro!", n)
    }
}
func conferirPrimo(n int) bool {
    for i := 2; i <= n / 2; i++{
        if n % i == 0 {
            return false
            break
        }
    }
    return true
}
```
