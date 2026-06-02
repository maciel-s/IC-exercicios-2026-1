package main

import (
	"fmt"
)

func main() {
	var cont, acum, tamArray int
	var endMaior, endMenor int
	var media float32
	var arrayIC []int

	tamArray = 20
	arrayIC = make([]int, tamArray)

	iniciaArrayShuffle(arrayIC)
	imprimir(arrayIC)

	endMaior = 0
	endMenor = 0
	acum = arrayIC[0]

	for cont = 1; cont < tamArray; cont++ {
		acum += arrayIC[cont]

		if arrayIC[cont] > arrayIC[endMaior] {
			endMaior = cont
		} else {
			if arrayIC[cont] < arrayIC[endMenor] {
				endMenor = cont
			}
		}
	}

	media = float32(acum) / float32(tamArray)

	fmt.Println("O maior numero eh: ", arrayIC[endMaior])
	fmt.Println("O menor numero eh: ", arrayIC[endMenor])
	fmt.Println("A media eh: ", media)

	troca(endMaior, endMenor, arrayIC)

	imprimir(arrayIC)
}
