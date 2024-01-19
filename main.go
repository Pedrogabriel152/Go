package main

import (
	"fmt"
	"time"
)

// type Pessoa struct {
// 	Nome      string
// 	Sobrenome string
// 	Idade     uint8
// 	Status    bool
// }

// func (p Pessoa) String() string {
// 	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d ano.", p.Nome, p.Idade)
// }

// type Categoria struct {
// 	Nome string
// 	Pai  *Categoria
// }

// func (c Categoria) HasParent() bool {
// 	return c.Pai != nil
// }

// func (c *Categoria) SetPai(pai *Categoria) {
// 	c.Pai = pai
// }

// type Document interface {
// 	Doc() string
// }

// type Pessoa struct {
// 	Nome   string
// 	Idade  int
// 	Status bool
// }

// type PessoaFisica struct {
// 	Pessoa
// 	Sobrenome string
// 	cpf       string
// }

// type PessoaJuridica struct {
// 	Pessoa
// 	RazaoSocial string
// 	Cnpj        string
// }

// func (p PessoaFisica) String() string {
// 	return fmt.Sprintf("Olá meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
// }

// func (pf PessoaFisica) Doc() string {
// 	return fmt.Sprintf("Meu cpf é %s.", pf.cpf)
// }

// func (pf PessoaJuridica) Doc() string {
// 	return fmt.Sprintf("Meu cnpj é %s.", pf.Cnpj)
// }

// func show(d Document) {
// 	if d, ok := d.(PessoaFisica); ok {
// 		fmt.Println((d.Sobrenome))
// 	}

// 	switch d.(type) {
// 	case PessoaFisica:
// 		fmt.Println((d.(PessoaFisica).Sobrenome))
// 	case PessoaJuridica:
// 		fmt.Println((d.(PessoaJuridica).RazaoSocial))
// 	}

// 	fmt.Println(d)
// 	fmt.Println(d.Doc())
// }

// func main() {
// 	// idades := make(map[string]uint8)

// 	// idades["Pedro"] = 21
// 	// idades["Gabi"] = 19

// 	// if val, ok := idades["Pedro"]; ok {
// 	// 	fmt.Println(val, ok)
// 	// }

// 	// p := Pessoa{
// 	// 	Nome:      "Pedro",
// 	// 	Sobrenome: "Gabriel",
// 	// 	Idade:     21,
// 	// 	Status:    true,
// 	// }

// 	// cat := Categoria{
// 	// 	Nome: "Categoria 1",
// 	// }
// 	// cat.SetPai(&Categoria{Nome: "Pai"})

// 	// if !cat.HasParent() {
// 	// 	fmt.Println("Não tem pai")
// 	// }

// 	// fmt.Println(p)

// 	p := PessoaFisica{
// 		Pessoa{
// 			Nome:   "Pedro",
// 			Idade:  21,
// 			Status: true,
// 		}, "Gabriel", "70683499190",
// 	}
// 	show(p)
// 	fmt.Println(p)
// }

func numeros(n chan<- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("\nEscrito doo chanel: %d", i)
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Printf("\nFim da escrita:")
	close(n)
}

func letras(done chan<- bool) {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)

		time.Sleep(time.Millisecond * 230)
	}

	done <- true
}

func main() {
	cn := make(chan int, 3)
	// cl := make(chan bool)

	go numeros(cn)
	// go letras(cl)

	for v := range cn {
		fmt.Printf("\nLido doo chanel: %d", v)
		time.Sleep(time.Millisecond * 230)
	}

	// <-cn
	// <-cl
	fmt.Println("\nFim da execução!")
}
