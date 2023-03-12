package main

import "fmt"

func main() {
	var v interface{} = 2.12
	var x interface{} = 10
	var y interface{} = "Hello, world!"
	var z interface{} = true

	Show(v)
	Show(x)
	Show(y)
	Show(z)

}

func Show(t interface{}) {
	fmt.Printf("O tipo da variável é: %T e o valor é %v\n", t, t)
}
