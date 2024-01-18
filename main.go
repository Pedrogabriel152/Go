package main

import "fmt"

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

type Pessoa struct {
	Nome   string
	Idade  int
	Status bool
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	Cnpj        string
}

func (p PessoaFisica) String() string {
	return fmt.Sprintf("Olá meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

func main() {
	// idades := make(map[string]uint8)

	// idades["Pedro"] = 21
	// idades["Gabi"] = 19

	// if val, ok := idades["Pedro"]; ok {
	// 	fmt.Println(val, ok)
	// }

	// p := Pessoa{
	// 	Nome:      "Pedro",
	// 	Sobrenome: "Gabriel",
	// 	Idade:     21,
	// 	Status:    true,
	// }

	// cat := Categoria{
	// 	Nome: "Categoria 1",
	// }
	// cat.SetPai(&Categoria{Nome: "Pai"})

	// if !cat.HasParent() {
	// 	fmt.Println("Não tem pai")
	// }

	// fmt.Println(p)

	p := PessoaFisica{
		Pessoa{
			Nome:   "Pedro",
			Idade:  21,
			Status: true,
		}, "Gabriel", "70683499190",
	}

	fmt.Println(p)
}
