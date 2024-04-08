# Quesões Médias

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

