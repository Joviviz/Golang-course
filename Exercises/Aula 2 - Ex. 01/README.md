# Aula 2 - Exercicio 1

Leia 2 valores com uma casa decimal (x e y), que devem representar as coordenadas de um ponto em um plano. A seguir, determine qual o quadrante ao qual pertence o ponto, ou se está sobre um dos eixos cartesianos ou na origem  (x = y = 0). Se o ponto estiver na origem, escreva a mensagem “Origem”. Se o ponto estiver sobre um dos eixos escreva “Eixo X” ou “Eixo Y”, conforme for a situação.

```go
package main

import "fmt"

func main() {
    var x, y float64
    
    fmt.Print("Digite a coordenada x: ")
    fmt.Scan(&x)
    fmt.Print("Digite a coordenada y: ")
    fmt.Scan(&y)
    
    check_coords(x, y)
}

func check_coords(x float64, y float64){
    if x == 0 && y == 0 {
        fmt.Print("Origem")
    } else if x == 0 {
        fmt.Print("Eixo Y")
    } else if y == 0 {
        fmt.Print("Eixo X")
    } else if x > 0 {
        if y > 0 {
            fmt.Print("Primeiro Quadrante")
        } else {
            fmt.Print("Quarto Quadrante")
        }
    } else {
        if y > 0 {
            fmt.Print("Segundo Quadrante")
        } else {
            fmt.Print("Terceiro Quadrante")
        }
    }
}
```
