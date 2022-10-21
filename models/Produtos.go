package models

import (
	"banco/db"
)

type Produto struct {
	Id         int
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

		p.Id = id
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

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao,preco, quantidade) values (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.Conecta()

	deletarProduto, err := db.Prepare("delete from produtos where id = ?")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)

	defer db.Close()

}

func EditaProduto(id string) Produto {
	db := db.Conecta()

	produtoDoBanco, err := db.Query("select * from produtos where id = ?", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id int
		var nome, descricao string
		var preco float64
		var quantidade int64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade

	}
	defer db.Close()
	return produtoParaAtualizar

}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.Conecta()

	AtualizaProduto, err := db.Prepare("update produtos set nome=?, descricao=?, preco=?, quantidade=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
