package main

import (
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	codigo, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	print(string(codigo))
	defer req.Body.Close()
}
