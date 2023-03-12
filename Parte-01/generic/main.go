package main

type Inteiro int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	x1 := map[string]int{"A": 2, "b": 5, "C": 8}
	x2 := map[string]float64{"A": 2.5, "b": 3.5, "C": 6.0}
	x3 := map[string]Inteiro{"A": 2, "b": 5, "C": 8}

	print(Soma(x1))
	print(Soma(x2))
	print(Soma(x3))
	println(Compara(10, 100))
}
