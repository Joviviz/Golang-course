# Questões Médias

# Q1
```go
package main

import "fmt"

func encontrarNumero(numeros []int) int {
	regIndex := 0

	for i := 0; i < len(numeros)-1; i++ {
		if numeros[i] != numeros[i+1] {
			regIndex = numeros[i]
		}
	}
	return regIndex
}

func main() {

	fmt.Println("Lista com os numeros : 1, 1, 1, 4, 0, 0")
	numeros := []int{1, 1, 1, 4, 0, 0}
	println("O número que não se repete na lista é: ", encontrarNumero(numeros))
}
```
# Q7
```go
package main

import "fmt"
import "math"

func main() {

	var x float64 = math.Cbrt(27)
	fmt.Println(x)
}

```
# Q9
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

# Q11
```go
package main

import "fmt"

func conferirTriangulo(a, b, c float64) {
	if a < b+c && b < a+c && c < b+a {
		println("É possível fazer um triângulo")
	} else {
		println("Não é possível fazer um triângulo")
	}

}

func main() {
	var a, b, c float64
	fmt.Println("Escreva um valor para os lados do triângulo: \n")

	fmt.Println("\nLado A: ")
	fmt.Scan(&a)
	fmt.Println("\nLado B: ")
	fmt.Scan(&b)
	fmt.Println("\nLado C ")
	fmt.Scan(&c)

	conferirTriangulo(a, b, c)
}
```

# Q12
```go
package main

import "fmt"

func mdc(n1, n2 int) int {
	if n2 == 0 {
		return n1
	}
	return mdc(n2, n1%n2)
}

func main() {
	fmt.Println(mdc(60, 40))
}

```
# Q13
```go
package main

import "fmt"

func conferirQuadradoPerfeito(x int) {
	if x == 1 {
		println("É um quadrado perfeito")
	} else {
		var y int
		var conf int

		for y = 2; y <= x/2; y++ {
			if y*y == x {
				//println("É um quadrado perfeito")
				conf = 1
				break
			} else {
				//println("Não é um quadrado perfeito")
				continue
			}
		}
		if conf == 0 {
			println("Não é um quadrado perfeito")
		} else {
			println("É um quadrado perfeito")
		}

	}
}

func main() {
	var x int

	fmt.Println("Escreva um número inteiro: ")
	fmt.Scan(&x)

	conferirQuadradoPerfeito(x)
}
```

# Q15
```go
package main

import "fmt"

func main() {
	var numero1 int
	fmt.Scan(&numero1)
	for i := numero1; i > 0; i-- {
		fmt.Println(i)
	}
}


```
# Q16
```go
package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int
	fmt.Scan(&numero1)
	fmt.Scan(&numero2)
	var x1 int = 0
	var x2 int = 0
	for i := 1; i <= numero1; i++ {
		if numero1%i == 0 {
			x1 += i
		}
	}
	for i := 1; i <= numero2; i++ {
		if numero2%i == 0 {
			x2 += i
		}
	}
	fmt.Println("soma dos divisores do primeiro número:", x1, "\nsoma dos divisores do segundo número:", x2)
	if numero1 == x2 || numero2 == x1 {
		fmt.Println("Os números são amigos")
	} else {
		fmt.Println("Os números não são amigos")
	}
}
```

