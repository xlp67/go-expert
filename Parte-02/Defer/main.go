package main

func main() {
	//executa todos os comando, pula o que tem "defer" e executa por útimo

	defer println("Terceira execução")
	println("Primeira execução")
	println("Segunda execução")

}
