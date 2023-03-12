package main

import "fmt"

func main() {
	teste := map[string]int{"A": 10, "B": 20, "C": 30}

	fmt.Println(teste)
	delete(teste, "A")
	fmt.Println(teste)
	teste["A"] = 10
	fmt.Println(teste)

	for _, y := range teste {
		fmt.Printf("O valor é %v\n", y)
	}

	experiment := map[string]int{}
	experiment["Giropops"] = 69
	fmt.Println(experiment)
}
