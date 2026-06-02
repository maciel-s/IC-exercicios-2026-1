package recurssao

import "fmt"

func exponeciacao(expoente int, base int) int {
	var resposta int
	if expoente == 1 {
		resposta = base
		fmt.Println(resposta)
	} else {
		resposta = base * exponeciacao(expoente-1, base)
		fmt.Println(resposta)
	}
	return resposta
}
