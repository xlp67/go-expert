package main

import "fmt"

type Endereco struct {
	Numero     int
	Logradouro string
	Cidade     string
	Estado     string
}

type Cliente struct {
	Name   string
	Idade  int
	Peso   float32
	Status bool
	Endereco
}
type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func (c Cliente) Desativar() {
	c.Status = false
	fmt.Printf("O cliente %s foi desativado\n", c.Name)
}

func main() {

	thiago := Cliente{
		Name:   "Thiago",
		Idade:  29,
		Peso:   1.80,
		Status: true,
	}
	thiago.Numero = 231
	thiago.Logradouro = "Pass. Jari"
	thiago.Cidade = "Capanema"
	thiago.Estado = "Pará"

	Desativacao(thiago)
}
