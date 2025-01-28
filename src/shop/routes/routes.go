package routes

import (
	"net/http"
	"shop/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
