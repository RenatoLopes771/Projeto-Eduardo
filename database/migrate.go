package database

import (
	"github.com/RenatoLopes771/Projeto-Eduardo/models"
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {

	DB.AutoMigrate(&models.Usuario{})

}
