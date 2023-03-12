package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		//Args contém os argumentos da linha de comando, começando com o nome do programa.
		req, err := http.Get("http://www.viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Requisição %v não encontrada!\n", err)

			//Observe que o tempo de execução do Go grava no erro padrão para pânicos e travamentos
			//fechar o Stderr pode fazer com que essas mensagens sejam enviadas para outro lugar, talvez para um arquivo aberto posteriormente.
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Resposta %v não encontrada!\n", err)
		}
		var data ViaCEP

		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parse da resposta %v não encontrada!\n", err)
		}
		arquivo, err := os.Create("dados.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo!\n")
		}
		defer arquivo.Close()
		_, err = arquivo.WriteString(fmt.Sprintf("CEP: %v\nLogradouro: %v\nBairro: %v\nCidade: %v\nUF: %v\n", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf))
		fmt.Printf("CEP: %v\nLogradouro: %v\nBairro: %v\nCidade: %v\nUF: %v\n", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf)
	}

}
