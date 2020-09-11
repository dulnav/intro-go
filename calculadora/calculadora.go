package main

func main() {

	resultadoFuncion := suma(8, 2)
	println(resultadoFuncion)

	resultadoFuncion = multiplica(165, 57)
	println(resultadoFuncion)

}

func suma(primerNumero, segundoNumero int) int {
	resultado := primerNumero + segundoNumero

	return resultado
}

func multiplica(primerNumero, segundoNumero int) int {
	resultado := primerNumero * segundoNumero

	return resultado
}
