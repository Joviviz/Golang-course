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
