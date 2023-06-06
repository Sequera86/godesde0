package mapas

import (
	"fmt"
)

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  50,
		"Chivas":       38,
		"Boca Juniors": 42,
	}
	fmt.Println(campeonato)

	for equipo, puntuaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntuaje de %d \n", equipo, puntuaje)
	}

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntuaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
