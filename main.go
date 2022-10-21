package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int64
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul", Preco: 30, Quantidade: 5},
		{"Tênis", "Confortavel", 90, 3},
		{"Fone", "Muito Bom", 59, 10},
		{"Fone Nike", "Muito Bom", 150, 10},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
