package main
import "fmt"

func main() {
    var a, b, c float64
    
    fmt.Print("Digite o valor de A: ")
    fmt.Scan(&a)
    fmt.Print("Digite o valor de B: ")
    fmt.Scan(&b)
    fmt.Print("Digite o valor de C: ")
    fmt.Scan(&c)

    solver(a, b, c)
}

func solver(a float64, b float64, c float64){
    fmt.Print("\n A área do triângulo retângulo que tem A por base e C por altura = ", (a*c)/2 )
    
    fmt.Print("\n\n A área  do círculo de raio C. (pi = 3.14159) = ", 3.14159*c*c)
    
    fmt.Print("\n\n A área do trapézio que tem A e B por bases e C por altura  = ", (a+b)*c/2)

    fmt.Print("\n\n A área do quadrado que tem lado B  = ", b*b)
    
    fmt.Print("\n\n A área do retângulo que tem lados A e B.  = ", a*b)
        
}
