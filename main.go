package main

import (
	"fmt"
	"strconv"
)

func Hello(nome string) {
	fmt.Println("Hello", nome, "!")
}

func Sum(a, b int) int {
	return a + b
}

func ConvertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)

	if err != nil {
		return
	}

	total = a + i

	return
}

func main() {
	Hello("Pedro")
	fmt.Println("total:", Sum(10, 29))

	total, err := ConvertAndSum(10, "29")
	fmt.Println("total:", total, err)
}
