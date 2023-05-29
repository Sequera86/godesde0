package ejercicios

import (
	"strconv"
)

func Devuelve2valores(palabra string) (int, string) {
	numero, err := strconv.Atoi(palabra)
	if err != nil {
		texto := "Hubo un error"
		return 0, texto
	}
	if numero > 100 {
		texto := "Es mayor a 100"
		return numero, texto
	} else {
		texto := "Es menor a 100"
		return numero, texto
	}

}
