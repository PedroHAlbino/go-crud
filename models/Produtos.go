package models

import (
	"banco/db"
	"log"
)

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int64
}

func BuscaTodosOsProdutos() []Produto {
	db := db.Conecta()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id int
		var nome, descricao string
		var preco int
		var quantidade int

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = float64(preco)
		p.Quantidade = int64(quantidade)

		produtos = append(produtos, p)

	}
	defer db.Close()
	return produtos

}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.Conecta()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao,preco, quantidade) values($1,$2,$3,$4")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	log.Println(insereDadosNoBanco)

	defer db.Close()
}
