package main
import "fmt"

func main() {
    var n1, n2 float64
    
    fmt.Println("Digite o primeiro numero: ")
    fmt.Scan(&n1)
    fmt.Println("Digite o segundo numero: ")
    fmt.Scan(&n2)
    
    soma(&n1, &n2)
}

func soma(n1,n2 *float64){
    fmt.Printf("%.2f + %.2f = %.2f", *n1, *n2, *n1+*n2)
}
