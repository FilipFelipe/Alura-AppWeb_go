package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"text/template"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	Port := os.Getenv("PORT")
	fmt.Println("Server: http://localhost:" + Port)
	http.HandleFunc("/", index)
	http.ListenAndServe(":"+Port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaBancoDeDados()
	todosProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for todosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = todosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}

func conectaBancoDeDados() *sql.DB {
	urlDB := os.Getenv("DB")
	connStr := urlDB
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Erro em abrir conexão com o banco de dados ", err)
		panic(err.Error())
	}
	return db
}
