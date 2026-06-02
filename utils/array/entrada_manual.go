package array

import "fmt"

func entradaManualArray(array []int) {
	var cont int
	var tamArray int

	tamArray = len(array)
	fmt.Println("Calculando estatisticas - Digite ", tamArray, " numeros, um por vez: ")

	for cont = 0; cont < tamArray; cont++ {
		fmt.Scan(&array[cont])
	}
}
