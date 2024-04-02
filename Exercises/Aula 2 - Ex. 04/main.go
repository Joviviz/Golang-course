package main
import "fmt"

func main() {
    var a, b, c, d int
    
    fmt.Print("Digite o valor de A: ")
    fmt.Scan(&a)
    fmt.Print("Digite o valor de B: ")
    fmt.Scan(&b)
    fmt.Print("Digite o valor de C: ")
    fmt.Scan(&c)
    fmt.Print("Digite o valor de D: ")
    fmt.Scan(&d)

    diferenca := (a * b) - (c * d)
    fmt.Println("\nDiferença do produto de C e D : ", diferenca)
}
