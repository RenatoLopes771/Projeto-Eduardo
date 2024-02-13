package controllers

import (
	"github.com/RenatoLopes771/Projeto-Eduardo/database"
	"github.com/RenatoLopes771/Projeto-Eduardo/models"
	"github.com/RenatoLopes771/Projeto-Eduardo/util"
	"github.com/gin-gonic/gin"
)

func usuarioReadByEmail(email string) (bool, models.Usuario) {
	usuario := models.Usuario{}

	DB := database.Connect()

	dbResult := DB.Where("email = ?", email).First(&usuario)

	if dbResult.Error != nil || dbResult.RowsAffected == 0 {
		return false, usuario
	}

	return true, usuario
}

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

		c.JSON(400, gin.H{
			"Message": "Email invalido",
			"Nome":    body.Email,
		})
		return
	}

	usuarioExiste, _ := usuarioReadByEmail(body.Email)

	if usuarioExiste {
		c.JSON(400, gin.H{
			"Message": "Email ja cadastrado",
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

// @Description Passe um email de um usuario existente. Preencha somente os campos que quer atualizar.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param nome path string false "Nome novo"
// @Param emailAtual path string true "E-mail atual"
// @Param emailNovo path string false "E-mail novo"
// @Param senhaAtual path string true "Senha atual"
// @Param senhaNova path string false "Senha nova"
// @Router /usuarios/atualizar [put]
func UsuariosUpdate(c *gin.Context) {

	var body struct {
		Nome       string
		EmailAtual string
		EmailNovo  string
		SenhaAtual string
		SenhaNova  string
	}

	c.BindJSON(&body)

	usuarioExiste, usuarioAntigo := usuarioReadByEmail(body.EmailAtual)

	if !usuarioExiste {
		c.JSON(400, gin.H{
			"Message": "Usuario não existe no banco de dados",
			"E-mail":  body.EmailAtual,
		})
		return
	}

	if body.SenhaAtual != usuarioAntigo.Senha {
		c.JSON(400, gin.H{
			"Message": "Senha incorreta",
			"Nome":    body.EmailAtual,
		})
		return
	}

	var nome, email, senha string
	var mensagem []string

	if body.Nome == "" {
		mensagem = append(mensagem, "Nome não atualizado")
		nome = usuarioAntigo.Nome
	} else {
		nome = body.Nome
	}

	if body.EmailNovo == "" {
		mensagem = append(mensagem, "Email não atualizado")
		email = usuarioAntigo.Email
	} else if !util.ValidarEmail(body.EmailNovo) {
		mensagem = append(mensagem, "Email não atualizado, o informado é inválido")
		email = usuarioAntigo.Email
	} else {
		email = body.EmailNovo
	}

	if body.SenhaNova == "" {
		mensagem = append(mensagem, "Senha não atualizada")
		senha = usuarioAntigo.Senha
	} else {
		senha = body.SenhaNova
	}

	DB := database.Connect()

	dbResult := DB.Model(&usuarioAntigo).Updates(models.Usuario{
		Nome:  nome,
		Email: email,
		Senha: senha,
	})

	if dbResult.Error != nil {
		c.JSON(400, gin.H{
			"Message": "Erro ao atualizar usuario",
			"Usuario": usuarioAntigo,
		})
		return
	}

	c.JSON(200, gin.H{
		"Message":     "Usuario atualizado com sucesso",
		"Usuario":     usuarioAntigo,
		"Observações": mensagem,
	})
}
