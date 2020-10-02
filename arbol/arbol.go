package main

func main() {
	construirArbolDeNavidad()
}

func construirArbolDeNavidad() {

	altura := 5

	longitudEspacios := calcularLongitud(altura)

	imprimirLinea(longitudEspacios, 1, `*`)

	for i := 1; i <= altura; i++ {
		longitudCaracteres := calcularLongitud(i)

		imprimirLinea(longitudEspacios-longitudCaracteres/2, longitudCaracteres, "^")
	}

	imprimirLinea(longitudEspacios, 1, `"`)

}

func imprimirLinea(cantidadEspacios int, cantidadCaracteres int, caracter string) {
	espacios := ""
	caracteres := ""

	for i := 0; i < cantidadEspacios; i++ {
		espacios += " "
	}

	for i := 0; i < cantidadCaracteres; i++ {
		caracteres += caracter
	}

	println(espacios, caracteres)
}

func calcularLongitud(altura int) int {
	return altura*2 - 1
}
