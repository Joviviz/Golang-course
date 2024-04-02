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
