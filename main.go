package main

import (
	"github.com/RenatoLopes771/Projeto-Eduardo/database"
	_ "github.com/RenatoLopes771/Projeto-Eduardo/docs"
	"github.com/RenatoLopes771/Projeto-Eduardo/routes"
)

// @title Projeto do Desafio Go CRUD
// @version 1.0
// @description Para saber mais leia o README.md
// @host localhost:3429
// @BasePath /

func main() {

	// migrar o banco de dados
	database.Migrate()

	// servidor
	routes.DefineRoutes()

}
