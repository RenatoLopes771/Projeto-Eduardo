package database

import (
	"github.com/RenatoLopes771/Projeto-Eduardo/models"
)

func Migrate() {

	Connect().AutoMigrate(&models.Usuario{})

}
