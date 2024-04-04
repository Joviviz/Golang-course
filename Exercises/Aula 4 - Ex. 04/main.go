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
