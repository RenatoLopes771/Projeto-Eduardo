package main

import (
	"github.com/RenatoLopes771/Projeto-Eduardo/database"
	"github.com/RenatoLopes771/Projeto-Eduardo/initializers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// Sqlite driver based on CGO "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
)

func init() {
	initializers.LoadEnv()
}

var DB *gorm.DB

func main() {

	DB = database.Connect()

	database.Migrate(DB)

	// servidor
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pppp",
		})
	})

	r.Run() // http://localhost:3429
}
