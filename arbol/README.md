# Crea un arbolito de navidad con Go.

Al final del taller habrás aprendido sobre bucles y escrito un programa que genera un arbolito de navidad. ¡Empecemos!

## Bucles

Es muy común que cuando estés programando necesites repetir una operación muchas veces. Con las herramientas que hemos aprendido hasta ahora, si quisiéramos imprimir la palabra "hola" 5 veces tendríamos que escribirlo de la manera siguiente:

```go
    println("Hola")
    println("Hola")
    println("Hola")
    println("Hola")
    println("Hola")
```

¡Sigue este [enlace](https://play.golang.org/p/N2Z8GYC1ipH) para ir al Go playground y ver este código en acción!

Muchas veces cuando escribes programas tienes que repetir la misma acción miles de veces. Imagínate lo complicado que sería tener que repetir la misma linea de código tantas veces :weary:. ¡Afortunadamente tenemos los bucles al rescate!

La acción de repetir algo en informática, se llama iteración. En Go se representa de la siguiente manera:

```go
    for i := 1; i <= 5; i++ {
        println("Hola")
    }
```

La palabra "for" en inglés significa "para" así que cuando empezamos un bucle es como decir "**Para** cada elemento haz...".

¿Y qué significa `i := 1; i <= 5; i++`? :confounded:.

Recuerda que a las variables se les puede cambiar el valor que tienen. La primera acción es asignarle un valor de `1` a la variable `i` (`i := 1`), esta variable nos va a servir para guardar la cuenta de cuantas veces hemos realizado la acción que estamos repitiendo.

Ahora vamos a definir cuantas veces queremos que se repita la acción. Si queremos repetir la acción 5 veces, lo que queremos decir es: "Mientras `i` tenga un valor menor o igual a `5`". Esto se representa de la siguiente manera `i <= 5`.

¿Y `i++`? En Go si le quieres sumar 1 a una variable que contenga un numero sólamente tienes que agregarle un `++` a tu variable. Así que `i + 1` es lo mismo que `i++`.

¡Ahora ya aprendiste sobre bucles! Ve como funciona el código de arriba siguiendo este [enlace](https://play.golang.org/p/B4LNJHQLK6B).

![Imagen de un gophers corriendo](../assets/gophers_corriendo.png)
