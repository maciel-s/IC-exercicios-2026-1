package recurssao

import "fmt"

func fatorial(numero int) int {
	var resposta int
	fmt.Print(numero, " - ")

	if numero == 0 {
		resposta = 1
	} else {
		resposta = (numero * fatorial(numero-1))
	}

	fmt.Println(resposta)

	return resposta
}
