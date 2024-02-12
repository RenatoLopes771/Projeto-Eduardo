package database

import (
	"log"

	"gorm.io/driver/sqlite"
	// Sqlite driver based on CGO "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
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
