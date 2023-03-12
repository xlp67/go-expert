package main

import (
	"bufio"
	"os"
)

func main() {
	linebyline, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(linebyline)
	buffer := make([]byte, 2)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		println(string(buffer[:n]))
	}
	// ----------------------  CLOSE FILE   ----------------------
	linebyline.Close()
}
