package main

import "fmt"

const a = "Hello, world!"

type ID int

var (
	b        bool    = true
	c        int     = 10
	d        float64 = 3.14159
	e        ID      = 1
	meuarray [3]int
)

func main() {
	meuarray[0] = 10
	meuarray[1] = 20
	meuarray[2] = 30
	for x, y := range meuarray {
		fmt.Printf("O valor do indície %d é %d \n", x, y)
	}
}
