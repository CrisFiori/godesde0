package main

import (
	
	"fmt"
	
	"home/cfiori/go/src/github.com/CrisFiori/godesde0/ejercicios"
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
*/
	nombre, cantidadPersonas := ejercicios.ConviertoaEntero("102")
		fmt.Println(nombre)
		fmt.Println(cantidadPersonas)


}
 