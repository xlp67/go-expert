package main

import "fmt"

type A struct {
	Name string
}
type P struct {
	Name string
}

func (A) B() {
	fmt.Printf("Função B com a struct A\n")
}

func (P) B() {
	fmt.Printf("Função B com a struct P\n")
}

type C interface {
	B()
}

func D(c C) {
	c.B()
}

func maini() {
	aaa := A{
		Name: "Alan",
	}
	ppp := P{
		Name: "Beto",
	}

	D(aaa)
	fmt.Println(aaa.Name)
	D(ppp)
	fmt.Println(ppp.Name)

}
