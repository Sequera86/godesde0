package main

import (
	/*"github.com/Sequera86/godesde0/ejercicios"
	"github.com/Sequera86/godesde0/variables"
	*/
	"github.com/Sequera86/godesde0/teclado"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	numero, cadena := ejercicios.Devuelve2valores("500")
	fmt.Println(numero, cadena)
	*/
	teclado.IngresosNumeros()
}
