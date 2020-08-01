package main

// Agregare comentarios que expliquen al código.

func main() {

	resultado := suma(8, 2)
	println(resultado)

	resultado = resta(10, 5)
	println(resultado)

}

func suma(primerNumero, segundoNumero int) (resultado int) {
	resultado = primerNumero + segundoNumero

	return resultado
}

func resta(primerNumero, segundoNumero int) (resultado int) {
	resultado = primerNumero - segundoNumero

	return resultado
}
