package routes

import (
	"api-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.ExibePersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.ExibePersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
