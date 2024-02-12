package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// github.com/mattn/go-sqlite3

	DB, err := gorm.Open(sqlite.Open("banco.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Falha na conexao ao banco de dados")
	}

	return DB
}
