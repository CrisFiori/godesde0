package ejercicios

import (
	"strconv"
)


func ConviertoaEntero (personas string )(int, string){
	cantidadPersonas, err := strconv.Atoi(personas)
	if err != nil{
		return 0, "Hubo un error"+ err.Error()
	}
	if cantidadPersonas > 100 {
		return cantidadPersonas, "Es mayor a 100"
	}else {
		return cantidadPersonas,"Es menor a 100"
	}

}