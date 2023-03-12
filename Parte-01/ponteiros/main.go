package main

import "fmt"

type Conta struct {
	Saldo int
	Tipo  string
}

func (c *Conta) simular(valor int) int {
	c.Saldo += valor
	c.Tipo = "Poupança"
	println(c.Saldo)
	fmt.Printf("%v\n \n", c.Tipo)
	return c.Saldo
}

func main() {
	conta := Conta{
		Saldo: 100,
		Tipo:  "Corrente",
	}

	conta.simular(200)
	println(conta.Saldo)
	fmt.Printf("%v\n \n \n", conta.Tipo)

	a := 1
	b := &a
	fmt.Printf("%v \n", b)
	a = 3
	fmt.Printf("%v \n", *b)
	*b = 10
	fmt.Printf("%v \n", a)

}
