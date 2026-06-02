package main

import "fmt"

func main() {
	var cont, temp, acum int
	var media float64
	var endMaior, endMenor int
	var arrayIC []int

	arrayIC = make([]int, 5)

	fmt.Println("Calculando estatísticas - Digite 5 numeros, um por vez!")
	for cont = 0; cont < 5; cont++ {
		fmt.Scan(&arrayIC[cont])
	}

	fmt.Print("arrayIC[ ")
	for cont = 0; cont < 5; cont++ {
		fmt.Print(arrayIC[cont], " ")
	}
	fmt.Println("]")
	endMaior = 0
	endMenor = 0
	acum = arrayIC[0]
	for cont = 1; cont < 5; cont++ {
		acum += arrayIC[cont]
		if arrayIC[cont] > arrayIC[endMaior] {
			endMaior = cont
		} else {
			if arrayIC[cont] < arrayIC[endMenor] {
				endMenor = cont
			}
		}
	}

	media = float64(acum) / float64(cont)
	fmt.Printf("Media: %.2f\n", media)
	fmt.Printf("Maior valor: %d\n", arrayIC[endMaior])
	fmt.Printf("Menor valor: %d\n", arrayIC[endMenor])

	// trocando o maior com o menor
	temp = arrayIC[endMaior]
	arrayIC[endMaior] = arrayIC[endMenor]
	arrayIC[endMenor] = temp

	fmt.Print("arrayIC[ ")
	for cont = 0; cont < 5; cont++ {
		fmt.Print(arrayIC[cont], " ")
	}
	fmt.Println("]")
}
