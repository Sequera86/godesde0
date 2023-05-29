package ejercicios

import (
	"strconv"
)

func Devuelve2valores(palabra string) (int, string) {
	if numero, _ := strconv.Atoi(palabra); numero > 100 {
		texto := "Es mayor a 100"
		return numero, texto

	} else {
		texto := "Es menor a 100"
		return numero, texto
	}

}
