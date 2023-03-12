package main

import "fmt"

func main() {
	//Abaixo vemos o exemplo de uma func dentro de outra func
	xxt := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) * 2
	}()

	fmt.Println(xxt)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total = numero + total

	}
	return total
}
