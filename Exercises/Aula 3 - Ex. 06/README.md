# Aula 3 - Exercício 6

Faça um algoritmo que leia vários números inteiros e mostre o maior deles. A leitura deve ser interrompida quando for lido o valor zero.

```go
package main
import "fmt"

func main() {
    var n, maior_n int
    fmt.Println("Escreva uma sequência de números inteiros positivos. Digite 0 para parar dar fim a sequência.")
    
    pablo(n, maior_n)
    
}

func pablo(n, maior_n int){
        for {
            fmt.Print("Digite um numero: ")
            fmt.Scan(&q)
            
            if n == 0{
                break
            }
        
            if n > maior_n{
                maior_n = n
            }
        }
}
```
