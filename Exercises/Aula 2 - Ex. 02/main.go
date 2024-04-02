package main

import "fmt"

func main() {
    var n1, n2 float64
    var operation int
    
    fmt.Print("Digite o primeiro número: ")
    fmt.Scan(&n1)
    fmt.Print("Digite o segundo número: ")
    fmt.Scan(&n2)
    
    fmt.Print("Escolha o código de operação: \n 1 - Soma \n 2 - Subtração \n 3 - Divisão \n 4 - Multiplicação \n")
    fmt.Scan(&operation)
    
    calculator(n1, n2, operation)
}

func calculator(n1 float64, n2 float64, operation int){
    if operation == 1 {
        fmt.Println("Soma = ", n1 + n2)
    } else if operation == 2 {
        fmt.Println("Subtração = ", n1 - n2)
    } else if operation == 3 {
        fmt.Println("Divisão = ", n1 / n2)
    } else if operation == 4 {
        fmt.Println("Multiplicação = ", n1 * n2)
    } else {
        fmt.Print("Valor de operação inválido...")
    }
}
