package main

// Voy a limpiar este código para que sea mas fácil de entender y
// le agregaré una estrellita hasta arriba :)

func main() {
	imprimirArbolDeNavidad(5)
}

func imprimirArbolDeNavidad(altura int) {
	if altura < 1 {
		return
	}

	ancho := altura*2 - 1

	for i := 1; i <= altura; i++ {
		salida := 0
		if i == 1 {
			salida = 1

		} else {
			salida = i*2 - 1

		}

		construirArbol(ancho-salida/2, salida, "o")

	}

	construirArbol(ancho, 1, `"`)

}

func construirArbol(espacio int, cuenta int, caracter string) {
	imprimeEspacio := ""
	imprimeCaracter := ""

	for i := 0; i < espacio; i++ {
		imprimeEspacio += " "
	}

	for i := 0; i < cuenta; i++ {
		imprimeCaracter += caracter
	}

	println(imprimeEspacio, imprimeCaracter)
}
