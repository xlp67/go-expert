package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		request, err := http.Get("https://" + url)
		if err != nil {
			panic(err)
		}
		defer request.Body.Close()
		result, err := io.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}

		print(string(result))

	}
}
