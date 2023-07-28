package main

import (
	//"fmt"
	//"github.com/Sequera86/godesde0/ejercicios"
	//"github.com/Sequera86/godesde0/files"
	//"github.com/Sequera86/godesde0/variables"
	//"github.com/Sequera86/godesde0/teclado"
	//"github.com/Sequera86/godesde0/funciones"
	//"github.com/Sequera86/godesde0/arreglos_slices"
	//"github.com/Sequera86/godesde0/users"
	e "github.com/Sequera86/godesde0/ejerinterfaces"
	"github.com/Sequera86/godesde0/modelos"
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
	//teclado.IngresosNumeros()
	//fmt.Println(ejercicios.Multipilicacion())
	//files.Sumatabla()
	//files.LeoArchivo2()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponencia(2)
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlice()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()
	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)
}
