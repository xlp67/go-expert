package main

func main() {
	a := 4
	b := 2

	if a > b {
		print("x\n")
	} else {
		print("y\n")
	}
	// "ou" >> "||"
	// "e" >> "&&"

	switch a {
	case 1:
		print("a\n")
	case 2:
		print("b\n")
	case 3:
		print("c\n")
	default:
		print("Não é nenhum dos valores\n")
	}
}
