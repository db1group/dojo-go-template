package example

import "fmt"

var (
	externaGeral  string = "qualquer coisa"
	externaGeral2        = false
)

func privadaSemNada() {
	fmt.Println("Sem parametro e sem retorno")
}

func privadaComParametroEComRetorno(myparameter int) string {
	var complete string = "um valor qualquer"
	fmt.Println(myparameter)
	fmt.Println(complete)
	return "private example"
}

func PublicaComParametroEComRetorno(myparameter bool) string {
	var complete = "um valor qualquer"
	fmt.Println(myparameter)
	fmt.Println(complete)
	return "public example"
}

func MultiParametrosEUmRetorno(first float32, second float64) string {
	complete := "um valor qualquer"
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(complete)
	return "private example"
}

func MultiParametrosEMultiRetornos(first int, second string) (string, int) {
	fmt.Println(externaGeral)
	fmt.Println(externaGeral2)
	fmt.Println(first)
	fmt.Println(second)
	return "private example", 0
}
