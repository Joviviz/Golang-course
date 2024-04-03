package main
import "fmt"

func main() {
	fizzbuzz()
}

func fizzbuzz() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz \n")
		} else if i%3 == 0 {
			fmt.Print("Fizz \n")
		} else if i%5 == 0 {
			fmt.Print("Buzz \n")
		} else {
			fmt.Printf("%d \n", i)
		}
	}
}
