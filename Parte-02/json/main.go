package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{
		Numero: 1,
		Saldo:  100,
	}
	// Marshal pega o resultado e guarda pra mim
	resultado, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	print(string(resultado))

	// Encoder gera um output que o json pode ser lido por outro sistema
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	jsonpuro := []byte(`{"n":2,"s":300}`)
	var xcontax Conta
	err = json.Unmarshal(jsonpuro, &xcontax)
	if err != nil {
		panic(err)
	}
	print(xcontax.Saldo)
}
