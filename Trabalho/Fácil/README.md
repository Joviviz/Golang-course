# Questões Facil

# Q1
```go
package main

import "fmt"

func main() {
   var numero1 int
   fmt.Scan(&numero1)
   var numero2 int
   fmt.Scan(&numero2)
   fmt.Println(numero1 + numero2)
}
```



# Q2
```go
package main

import "fmt"

func main() {
   var numero1 int
   fmt.Scan(&numero1)
   var numero2 int
   fmt.Scan(&numero2)
   var numero3 int
   fmt.Scan(&numero3)
   fmt.Println("a média é:", (numero3+numero2+numero1)/3)
}
```



# Q3
```go
package main

import "fmt"

func main() {
	var number int
	fmt.Print("Escreva um número: ")
	fmt.Scan(&number)

	fmt.Print("Método If/Else: ")
	even_odd_ifelse(number)

	fmt.Print("Método Switch: ")
	even_odd_switch(number)
}

func even_odd_ifelse(n int) {
	if n%2 == 0 {
		fmt.Print("Par\n")
	} else {
		fmt.Print("Impar\n")
	}
}

func even_odd_switch(n int) {
	switch {
	case n%2 == 0:
		fmt.Print("Par\n")

	default:
		fmt.Print("Impar\n")
	}
}

```
# Q4
```go
package main

import "fmt"

func main() {
	var numero1 int
	fmt.Scan(&numero1)
	var x int = 1
	for i := 0; i < numero1; i++ {
		x += x * i
	}
	fmt.Println(x)
}

```
# Q5
```go
package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Println("Informe um número")
	fmt.Scan(&x)
	fmt.Println("Informe um outro número")
	fmt.Scan(&y)

	if x >= y {
		fmt.Println(x, "é o maior número")

	} else {
		fmt.Println(y, "é o maior número")

	}

}
```
# Q6
```go
package main

import "fmt"

func main() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}


```
# Q7
```go

```
# Q8
```go
package main

import "fmt"

func main() {
	var numero1 int
	fmt.Scan(&numero1)
	fmt.Println("Fahrenheit:", (numero1*9/5)+32)
}


```
# Q9
```go
package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	var n3 int

	fmt.Println("Insira um número")
	fmt.Scan(&n1)

	fmt.Println("Insira um outro número")
	fmt.Scan(&n2)

	fmt.Println("Insira mais outro número")
	fmt.Scan(&n3)

	maior := n1

	if n2 > maior {
		maior = n2
	}

	if n3 > maior {
		maior = n3
	}

	fmt.Println(maior, "é o maior número")

}
```
# Q10
```go
package main

import "fmt"

func main() {
	var numero1 int
	fmt.Scan(&numero1)
	for i := 1; i <= 10; i++ {
		fmt.Println(numero1 * i)
	}
}


```
# Q11
```go
package main

import "fmt"

func main() {
	var numero1 int
	var x int = 0
	fmt.Scan(&numero1)
	for i := 2; i < numero1; i++ {
		if numero1%i == 0 {
			x += 1
		}
	}
	if x > 0 {
		fmt.Println("não é primo")
	} else {
		fmt.Println("é primo")
	}
}


```
# Q12
```go
package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int
	var numero3 int
	var peso1 int
	var peso2 int
	var peso3 int
	fmt.Println("Numero 1:")
	fmt.Scan(&numero1)
	fmt.Println("Numero 2:")
	fmt.Scan(&numero2)
	fmt.Println("Numero 3:")
	fmt.Scan(&numero3)
	fmt.Println("Peso 1:")
	fmt.Scan(&peso1)
	fmt.Println("Peso 2:")
	fmt.Scan(&peso2)
	fmt.Println("Peso 3:")
	fmt.Scan(&peso3)
	fmt.Println("A média ponderada é:", (numero1*peso1+numero2*peso2+numero3*peso3)/3)
}
```
# Q13
```go
package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int
	var numero3 int
	fmt.Println("Valor inicial")
	fmt.Scan(&numero1)
	fmt.Println("Taxa diária")
	fmt.Scan(&numero2)
	fmt.Println("Dias de atraso")
	fmt.Scan(&numero3)
	fmt.Println(numero1 + (numero2 * numero3))
}

```
# Q14
```go
package main

import "fmt"

func main() {
	var palavra string
	fmt.Scan(&palavra)
	var t int
	for _, value := range palavra {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			t++
		}
	}
	fmt.Println(t, "vogais")
}

```
# Q15
```go
package main

import "fmt"

func main() {
	var numero1 int
	var x int
	fmt.Scan(&numero1)
	for i := 1; i < numero1; i++ {
		if numero1%i == 0 {
			x += i
		}
	}
	if x == numero1 {
		fmt.Println("é número perfeito")
	} else {
		fmt.Println("não é número perfeito")
	}

}


```
# Q16
```go

```
# Q17
```go

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

```
# Q18
```go

```
# Q19
```go
package main

import "fmt"

func main() {
	var numero1 float64
	fmt.Scan(&numero1)
	fmt.Printf("O volume da esfera é:%.2f", 4/3*3.14*numero1)

}


```
# Q20
```go
package main

import "fmt"

func main() {
	var numero1 int
	fmt.Scan(&numero1)
	if numero1%4 == 0 && numero1%10 != 0 {
		fmt.Println("é bissexto")
	} else {
		fmt.Println("não é bissexto")
	}
}


```

