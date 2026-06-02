package main

import (
	"fmt"
	"ic_exercicios/utils/array"
)

func main() {
	var cont, acum, tamArray int
	var endMaior, endMenor int
	var media float32
	var arrayIC []int

	tamArray = 20
	arrayIC = make([]int, tamArray)

	array.IniciaArrayShuffle(arrayIC)
	array.Imprimir(arrayIC)

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

	array.Troca(endMaior, endMenor, arrayIC)

	array.Imprimir(arrayIC)
}
