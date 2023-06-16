package main

import (
	"home/cfiori/go/src/github.com/CrisFiori/godesde0/middlewares"
	//	"fmt"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)


	if os := runtime.GOOS; os =="linux" || os=="OS X."{
		fmt.Println("Esto no es Windows, es", os)
	}else {
		fmt.Println("Esto es Windows")
	}

	switch os:= runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("$s \n", os)
	}

	nombre, cantidadPersonas := ejercicios.ConviertoaEntero("102")
		fmt.Println(nombre)
		fmt.Println(cantidadPersonas)

	teclado.IngresoNumeros()
	iteraciones.Interar()
	fmt.Println(ejercicios.Multiplicando())

	files.SumaTabla()

	files.LeoArchivo()

	funciones.LlamarClosure()
	arreglos_slices.Capacidad()

	users.AltaUsuario()
	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)
	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)

	defer_panic.EjemploPanic ()
	canal1 := make(chan bool)
	go goroutines.MiNombreLento("Cristian Fiori", canal1)

	fmt.Println("Estoy aqui :")
	<-canal1
	webserver.MiWebServer()*/

	middlewares.MiMiddleware()

}
