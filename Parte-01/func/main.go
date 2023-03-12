package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := soma(1, 1)
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		fmt.Println(valor)
	}

}
func soma(a int, b int) (int, error) {
	if a+b >= 10 {
		return 0, errors.New("A soma dos valores resultam em um valor maior que 10")
	}

	return a + b, nil
}
