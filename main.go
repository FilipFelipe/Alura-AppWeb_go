package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/filipfelipe/Alura-AppWeb_go/routes"
	_ "github.com/lib/pq"
)

func main() {
	Port := os.Getenv("PORT")
	fmt.Println("Server: http://localhost:" + Port)
	routes.CarregaRotas()
	http.ListenAndServe(":"+Port, nil)
}
