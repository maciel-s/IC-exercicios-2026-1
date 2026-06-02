package main

import "math/rand"

func iniciaArrayShuffle(array []int) {
	var cont int
	var index1 int
	var index2 int
	var tamArray int

	tamArray = len(array)

	for cont = 0; cont < tamArray; cont++ {
		array[cont] = cont
	}

	for cont = 0; cont < tamArray; cont++ {
		index1 = rand.Intn(tamArray)
		index2 = rand.Intn(tamArray)

		troca(index1, index2, array)
	}
}
