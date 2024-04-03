package main

import "fmt"

func main() {
	fmt.Print("Numeros impares de 1 a 19: \n")
	for i := 1; i < 20; i++ {
		if i%2 != 0 {
			fmt.Printf("%d \n", i)
		}
	}
}
