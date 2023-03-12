package matematica

func Soma[T int | float64](x, y T) T {
	return x + y

}

type Carro struct {
	Marca string
}

func (Carro) Andar() string {
	return "carro andando!!!"
}
