package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func Multipilicacion() {
	for {
		fmt.Println("Ingrese numero a multiplicar : ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado no es un numero")
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i <= 10; i++ {
		println(numero1, " x ", i, " = ", numero1*i)
	}
}
