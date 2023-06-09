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
	for _, url := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + url + "/json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		defer req.Body.Close()
		resposta, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		var data ViaCEP
		err = json.Unmarshal(resposta, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		fmt.Println(data.Localidade)
	}
}
