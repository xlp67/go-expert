package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//---------------------- FILE CREATED ----------------------
	arquivo, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	//----------------------  WRITE FILE  ----------------------
	tamanho, err := arquivo.Write([]byte("Thiago Almeida"))
	if err != nil {
		panic(err)
	}
	//----------------------  READ FILE   ----------------------
	ler, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Size: %v\nContent: %v\n", tamanho, string(ler))
	//----------------------  OPEN FILE   ----------------------
	linebyline, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	//---------------------- BUFFER FILE  ----------------------
	reader := bufio.NewReader(linebyline)
	buffed := make([]byte, 1)
	for {
		n, err := reader.Read(buffed)
		if err != nil {
			break
		}
		println(string(buffed[:n]))
	}

	// ----------------------  CLOSE FILE   ----------------------
	linebyline.Close()
	// ---------------------- REMOVE FILE  ----------------------
	os.Remove("file.txt")
}
