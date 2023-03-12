package main

func main() {

	var MinhaVar interface{} = "Thiago William"

	res, ok := MinhaVar.(string)

	print(res, ok)
}
