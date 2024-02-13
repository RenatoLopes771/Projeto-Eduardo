package routes

import (
	"github.com/RenatoLopes771/Projeto-Eduardo/controllers"
	"github.com/RenatoLopes771/Projeto-Eduardo/initializers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnv()
}

func DefineRoutes() {

	r := gin.Default()

	// swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pppp",
		})
	})

	usuariosR := r.Group("usuarios")
	{
		usuariosR.POST("/criar", controllers.UsuariosCreate)

		usuariosR.GET("/", controllers.UsuariosReadAll)
	}

	r.Run()
}
