package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sequera86/godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func Grabatabla() {
	var texto string = ejercicios.Multipilicacion()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error)
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func Sumatabla() {
	var texto string = ejercicios.Multipilicacion()
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}
func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo", err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el writing del archivo", err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo", err.Error())
		return
	}
	fmt.Println(string(archivo))

}

func LeoArchivo2() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo", err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> ", registro)
	}
}
