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
