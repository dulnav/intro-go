package main

func main() {
	construirArbolDeNavidad()
}

func construirArbolDeNavidad() {

	altura := 5

	longitudEspacios := calcularLongitud(altura)

	imprimirLinea(longitudEspacios, 1, `*`)

	for i := 1; i <= altura; i++ {
		longitudCaracter := calcularLongitud(i)

		imprimirLinea(longitudEspacios-longitudCaracter/2, longitudCaracter, "^")

	}

	imprimirLinea(longitudEspacios, 1, `"`)

}

func imprimirLinea(espacio int, caracter int, elemento string) {
	espacios := ""
	caracteres := ""

	for i := 0; i < espacio; i++ {
		espacios += " "
	}

	for i := 0; i < caracter; i++ {
		caracteres += elemento
	}

	println(espacios, caracteres)
}

func calcularLongitud(altura int) int {
	return altura*2 - 1
}
