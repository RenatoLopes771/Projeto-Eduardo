package controllers

import (
	"github.com/RenatoLopes771/Projeto-Eduardo/database"
	_ "github.com/RenatoLopes771/Projeto-Eduardo/docs"
	"github.com/RenatoLopes771/Projeto-Eduardo/models"
	"github.com/RenatoLopes771/Projeto-Eduardo/util"
	"github.com/gin-gonic/gin"
)

// @Tags Usuarios
// @Accept json
// @Produce json
// @Param nome path string true "Nome"
// @Param email path string true "E-mail"
// @Param senha path string true "Senha"
// @Router /usuarios/criar [post]
func UsuariosCreate(c *gin.Context) {

	var body struct {
		Nome  string
		Email string
		Senha string
	}

	c.BindJSON(&body)

	if body.Nome == "" {
		c.JSON(400, gin.H{
			"Message": "Nome invalido",
			"Nome":    body.Nome,
		})
		return
	}

	DB := database.Connect()

	if body.Email == "" || !util.ValidarEmail(body.Email) {

		// TODO validacao se tem usuario com o mesmo email

		// DB.First(&models.Usuario{})

		c.JSON(400, gin.H{
			"Message": "Email invalido",
			"Nome":    body.Email,
		})
		return
	}

	if body.Senha == "" {
		c.JSON(400, gin.H{
			"Message": "Senha invalido",
			"Nome":    body.Senha,
		})
		return
	}

	usuario := models.Usuario{
		Nome:  body.Nome,
		Email: body.Email,
		Senha: body.Senha,
	}

	dbresult := DB.Create(&usuario)

	usuario.Senha = ""

	if dbresult.Error != nil {
		c.JSON(400, gin.H{
			"Message": "Erro na criacao de usuario",
			"Usuario": usuario,
		})
		return

	} else {
		c.JSON(200, gin.H{
			"Message": "Usuario criado",
			"Usuario": usuario,
		})
	}

}

// @Tags Usuarios
// @Produce json
// @Router /usuarios [get]
func UsuariosReadAll(c *gin.Context) {

	var usuarios []models.Usuario

	DB := database.Connect()

	dbresult := DB.Find(&usuarios)

	if dbresult.Error != nil {
		c.JSON(400, gin.H{
			"Message": "Erro na listagem de usuarios",
		})
		return
	}

	c.JSON(200, gin.H{
		"Usuarios": usuarios,
	})
}
