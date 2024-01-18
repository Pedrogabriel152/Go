package main

import (
	"fmt"
	"log"
	"os"
)

var (
	cara, coroa int
)

func lancarMoeda(lado string) {
	switch lado {
	case "cara":
		cara++
	case "coroa":
		coroa++

	default:
		fmt.Println("caiu em pé")
	}
}

func main() {
	a, b := 10, 13

	if a > b {
		fmt.Println("A é maior que B")
	} else if a < b {
		fmt.Println("B é maior A")
	} else {
		fmt.Println("B é iagaul a A")
	}

	switch {
	case a > b:
		fmt.Println("A é maior que B")
	case a < b:
		fmt.Println("B é maior A")
	default:
		fmt.Println("B é iagaul a A")
	}

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)

	_, err = file.Read(data)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))
}
