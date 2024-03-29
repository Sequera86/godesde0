package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func Multipilicacion() string {
	for {
		fmt.Println("Ingrese numero a multiplicar : ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado no es un numero")
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i <= 10; i++ {
		//println(numero1, " x ", i, " = ", numero*i)
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}
	return texto
}
