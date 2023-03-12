package main

import "fmt"

func main() {
	// Um slice trabalha com array por baixo dos panos e sempre que é dado um append,
	// ele cria outro array com a capacidade duplicada.
	// com isso é mais viável começar o slice próximo da capacidade desejada para evitar o processos extras.

	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	s = append(s, 110) //
	fmt.Printf("len=%d cap=%d %v\n", len(s[0:]), cap(s[0:]), s[0:])
}
