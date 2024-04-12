# Aula 2 - Exercicio Bonus - Fita Colorida 

> [!WARNING] 
> WORK IN PROGRESS

Roberto tem um conjunto de lápis com 10 tons diferentes de uma mesma cor, numerados de 0 a 9. Numa fita quadriculada, alguns quadrados foram coloridos inicialmente com o tom 0.

Roberto precisa determinar, para cada quadrado Q não colorido, qual é a distância dele para o quadrado mais próximo de tom 0.
A distância entre dois quadrados é definida com o número mínimo de movimentos para a esquerda, ou para a direita, para ir de um quadrado para o outro.
O quadrado Q, então, deve ser colorido com o tom cuja numeração corresponde à distância determinada. 
Se a distância for maior ou igual a 9, o quadrado deve ser colorido com o tom 9.

Seu programa deve colorir e imprimir a fita quadriculada dada na entrada.


## <ins>Entrada:<ins>

A primeira linha da entrada contém apenas um inteiro N, indicando o número de quadrados da fita. A segunda linha contém N números inteiros: “-1” se o quadrado não está colorido, e “0” se está colorido com o tom 0.

## <ins>Saída:<ins>

Seu programa deve escrever na saída a fita totalmente colorida, de acordo com a regra definida acima.

<br>

![Example img](https://cdn.discordapp.com/attachments/1224065809576366200/1224825235396628498/Captura_de_tela_2024-04-02_110631.png?ex=661ee664&is=660c7164&hm=aea90e02d726aaf9449783dfca39d947b58a2edc4a58574f202b84deaab6a2b7&)

<br>

```go
package main

import "fmt"

func definirFita(tamanho int) []int {

	fita := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		for {
			fmt.Print("\nRegistre uma tipagem de cor na posicao ", i+1)
			fmt.Print("\nNão colorido (-1) ou colorido com Tom 0 (0): \n")
			fmt.Scan(&fita[i])

			if fita[i] == -1 || fita[i] == 0 {
				break
			} else {
				fmt.Print("Valor invalido, tente novamente\n")
			}
		}
	}
	fmt.Print(fita)
	return fita
}

func abs(tomDeCor int) int {
	if tomDeCor < 0 {
		tomDeCor = tomDeCor * (-1)
	}
	return tomDeCor
}

func colorirFita(tamanho int, fita []int) {
	//fitaColorida := make([]int, tamanho)
	var tomDeCor int
	var tom0 int

	/*for i, cor := range fita {
		if cor == 0 {
			tom0 = i
		}
	}

	for i := 0; i < tamanho; i++ {
		if fita[i] == -1 {
			tomDeCor = i - tom0
			tomDeCor = abs(tomDeCor)
			fmt.Printf("%d", tomDeCor)
		} else {
			tomDeCor = 0
			fmt.Printf("%d", tomDeCor)
		}
	}*/

	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho; j++ {
			
		}
	}
}

func main() {
	var tamanho int
	fmt.Println("Tamanho da fita : ")
	fmt.Scan(&tamanho)

	fita := definirFita(tamanho)
	colorirFita(tamanho, fita)
}

```
```go
package main
import "fmt"

func definirFita(tamanho int) []int {

	fita := make([]int, tamanho)

	for i:= 0; i < tamanho; i++ {
		for {
			fmt.Println("\nRegistre uma tipagem de cor na posicao ", i+1)
			fmt.Println("Não colorido (-1) ou colorido com Tom 0 (0): ")
			fmt.Scan(&fita[i])

			if fita[i] == -1 || fita[i] == 0 {
				break
			} else {
				fmt.Print("Valor invalido, tente novamente\n")
			}
		}
	}
	fmt.Print(fita)
	return fita
}

func abs(tomDeCor int) int{
	if tomDeCor < 0{
		tomDeCor = tomDeCor* (-1)
	}
	return tomDeCor
}

func colorirFita(tamanho int,fita []int){
	fitaColorida := make([]int, tamanho)
	

	for i := 0; i < tamanho; i++ {
		if fita[i] == -1{
			distancia_0 := tamanho
			for j := 0; j < tamanho; j++ {
				if fita[j] == 0 && abs(j - i) < distancia_0{
					distancia_0 = abs(j - i)
				}
			}
			fitaColorida[i] = min(distancia_0, 9)
		}

	}

	print(fitaColorida)


}

func main() {
    var tamanho int
    fmt.Println("Tamanho da fita : ")
    fmt.Scan(&tamanho)

	fita := definirFita(tamanho)
	colorirFita(tamanho, fita)
}
```
