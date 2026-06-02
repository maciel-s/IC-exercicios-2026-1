package array

import "fmt"

func Imprimir(array []int) {
	var cont int
	var tamArray int

	tamArray = len(array)

	fmt.Print("arrayIC[ ")
	for cont = 0; cont < tamArray; cont++ {
		fmt.Print(array[cont], " ")
	}
	fmt.Println("]")

}
