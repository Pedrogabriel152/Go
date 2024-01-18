package main

import "fmt"

func main() {
	nomes := []string{"Pedro", "Daniel", "Thiago"}
	var i int
	for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for i < len(nomes) {
		fmt.Println(nomes[i])
		i++
	}
}
