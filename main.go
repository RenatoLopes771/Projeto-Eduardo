package main

import (
	"github.com/RenatoLopes771/Projeto-Eduardo/database"
	"github.com/RenatoLopes771/Projeto-Eduardo/routes"
)

func main() {

	// migrar o banco de dados
	database.Migrate()

	// servidor
	routes.DefineRoutes()

	// http://localhost:3429
}
