package array

import (
	"math/rand"
)

func fillArrayRandomico(array []int) {
	var cont int
	var valor int
	var tamArray int

	tamArray = len(array)

	for cont = 0; cont < tamArray; cont++ {
		valor = rand.Intn(tamArray)
		array[cont] = valor
	}

}
