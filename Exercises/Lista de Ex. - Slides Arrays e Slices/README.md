# Exercícios

1 - Crie um Array de inteiros com 5 elementos e imprima cada elemento em uma linha
separada. <br>
```go
package main

import "fmt"

func printArray(inteiros []int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Linha %d = %d \n", i+1, inteiros[i])
	}
}

func main() {
	var inteiros = [5]int{5, 4, 3, 2, 1}

	fmt.Printf("O Array é %d \n", inteiros)
	printArray(inteiros[:])

}
```


2 - Crie um Slice de strings vazio, adicione os nomes "João", "Maria" e "José" a ele e imprima o Slice. <br>
```go
package main

import "fmt"

func adicionarNomes(nomes []string) []string {
	fmt.Println("Adicionando nomes . . .")
	nomes = append(nomes, "João", "Maria", "José")
	return nomes
}

func main() {
	var nomes = []string{}

	fmt.Printf("Array nomes: %s \n", nomes)
	nomes = adicionarNomes(nomes)

	fmt.Printf("Array nomes: %s \n", nomes)

}
```

3 - Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante. <br>
```go
package main

import "fmt"

func removerTerceiro(nums []int) []int {
	fmt.Println("Removendo o terceiro . . .")
	nums = append(nums[:2], nums[3:]...)

	return nums
}

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice = %d \n", nums)

	nums = removerTerceiro(nums)
	fmt.Printf("Slice = %d \n", nums)
}
```

4 - Crie um Array de floats com 6 elementos e calcule a média dos valores armazenados no Array. <br>
```go
package main
import "fmt"

func definirArray(nums []float64) []float64{
	var valor float64

	for i := 0; i < 6; i++ {
		fmt.Printf("Adicione um valor na posicao %d: ", i+1)
		fmt.Scan(&valor)

		nums = append(nums, valor)
	}
	return nums
}

func calcularMedia(nums []float64) float64{
	somatorio := 0.0

	for i := 0; i < len(nums); i++ {
		somatorio += nums[i]
	}

	media := somatorio / float64(len(nums))
	return media
}

func main(){
	var nums []float64
	nums = definirArray(nums)

	fmt.Println("Sua lista : ")
	fmt.Println(nums)

	media := calcularMedia(nums)
	fmt.Printf("A média é : %.2f", media)

}
```

5 - Crie uma matriz bidimensional de inteiros com 3 linhas e 4 colunas. Inicialize cada elemento com o valor da soma do índice da linha e o índice da coluna. Imprima a matriz resultante. <br>
```go
package main
import "fmt"

func determinarMatriz(matriz [3][4]int) [3][4]int{
	for linha := range matriz{
		for coluna := range matriz[linha]{
			matriz[linha][coluna] = linha + coluna
			fmt.Println("Elemento[", linha, "][", coluna, "] = ", matriz[linha][coluna])
		}
	}

	return matriz
}

func main(){
	matriz := [3][4]int {{},{},{}}
	matriz = determinarMatriz(matriz)

	fmt.Println("\nSua matriz : ")
	fmt.Println(matriz)
}
```
