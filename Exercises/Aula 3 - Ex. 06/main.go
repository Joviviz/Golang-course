package main
import "fmt"

func main() {
    var n, maior_n int
    fmt.Println("Escreva uma sequência de números inteiros positivos. Digite 0 para parar e dar fim à sequência.")
    
    pablo(&n, &maior_n)
    fmt.Printf("O maior número é: %d", maior_n)
}

func pablo(n, maior_n *int) {
    for {
        fmt.Print("Digite um número: ")
        fmt.Scan(n)
        
        if *n == 0 {
            break
        }
    
        if *n > *maior_n {
            *maior_n = *n
        }
    }
}
