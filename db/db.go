package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConectaBancoDeDados() *sql.DB {
	urlDB := os.Getenv("DB")
	connStr := urlDB
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Erro em abrir conexão com o banco de dados ", err)
		panic(err.Error())
	}
	return db
}
