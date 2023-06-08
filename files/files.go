package files

import(
	"home/cfiori/go/src/github.com/CrisFiori/godesde0/ejercicios"
	"os"
	//"bufio"
	//"ioutil"
	"fmt"
)

var filename string = "./files/txt/texto.txt"

func GrabaTabla(){
	var texto string = ejercicios.Multiplicando()
	archivo, err:= os.Create(filename)
	if err != nil{
		fmt.Println("Error al crear el archivo"+err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla(){
	var texto string = ejercicios.Multiplicando()
	if Append(filename, texto) == false {
		fmt.Println("Error al concatenar contenido")
	} 
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil{
		fmt.Println("Error durante el Append" +err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil{
		fmt.Println("Error durante el WriteString"+err.Error())
		return false
	}

	arch.Close()
	return true

}