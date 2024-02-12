package routes

import (
	"github.com/RenatoLopes771/Projeto-Eduardo/controllers"
	"github.com/RenatoLopes771/Projeto-Eduardo/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func DefineRoutes() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pppp",
		})
	})

	r.POST("/usuarios/criar", controllers.UsuariosCreate)

	r.GET("/usuarios", controllers.UsuariosReadAll)

	r.Run()
}
