package routes

import (
	"net/http"

	"github.com/filipfelipe/Alura-AppWeb_go/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
