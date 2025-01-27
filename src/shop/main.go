package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Camiseta de algodão branca", Preco: 10.00, Quantidade: 10},
		{Nome: "Calça", Descricao: "Calça jeans", Preco: 20.00, Quantidade: 20},
		{Nome: "Tênis", Descricao: "Preto", Preco: 30.00, Quantidade: 30},
		{Nome: "Boné", Descricao: "Aba curva", Preco: 40.00, Quantidade: 40},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
