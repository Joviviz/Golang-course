package main
import "fmt"

func main() {
    var f, c, valor float64
    var escolha int
    
    fmt.Println("Digite para conversao de : ")
    fmt.Println("1 - Celsius -> Fahrenheit")
    fmt.Println("2 - Fahrenheit -> Celsius")
    
    fmt.Scanln(&escolha)
    
    switch escolha{
        case 1:
            celsiusParaFahrenheit(valor, c)
        case 2:
            fahrenheitParaCelsius(valor, f)
        default:
            fmt.Println("Valor invalido.")
    }
}

func celsiusParaFahrenheit (valor , c float64){
    fmt.Println("Digite a temperatura em Celsius: ")
    fmt.Scanln(&c)
    
    valor = c * 1.8 + 38
    fmt.Printf("O valor em Fahrenheit = %.2f F", valor)
}

func fahrenheitParaCelsius (valor, f float64){
    fmt.Println("Digite a temperatura em Fahrenheit: ")
    fmt.Scanln(&f)
    
    valor = (f - 38) / 1.8
    fmt.Printf("O valor em Celsius = %.2f C", valor)
}
