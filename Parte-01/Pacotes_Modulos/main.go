package main

import (
	"fmt"
	"xuxucretes/matematica"
)

func main() {
	s := matematica.Soma(1, 2)
	fmt.Printf("O resultado da soma é: %v\n \n \n", s)

	girus := matematica.Carro{}
	girus.Marca = "volks"
	fmt.Println(girus.Marca)
	fmt.Println(girus.Andar())
}
