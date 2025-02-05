package main

import (
	"api-rest/database"
	"api-rest/routes"
)

func main() {
	// models.Personalidades = []models.Personalidade{
	// 	{Id: 1, Nome: "Albert Einstein", Historia: "Físico alemão"},
	// 	{Id: 2, Nome: "Isaac Newton", Historia: "Físico inglês"},
	// }

	database.ConectaBancoDeDados()

	routes.HandleRequest()
}
