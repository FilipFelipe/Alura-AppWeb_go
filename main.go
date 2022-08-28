package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fmt.Println("Server: http://localhost:8000")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camisa", "Azul", 39, 1},
		{"Camisa", "Vermelho ", 59.55, 5},
		{"Camisa", "Vermelho ", 102.41, 2},
		{"Camisa", "Vermelho ", 200, 3},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
