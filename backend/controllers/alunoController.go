package controllers

import (
	"log"
	"net/http"
	"provaSuficiencia/config"
	"provaSuficiencia/models"

	"github.com/gin-gonic/gin"
)

//Controller de cada função;

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	config.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func CreateAluno(c *gin.Context) {
	var input models.Aluno
	// Faz o bind do json para o input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Cria um novo aluno no banco
	aluno := models.Aluno{Nome: input.Nome, Telefone: input.Telefone}
	config.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
func CreateDetailsAluno(c *gin.Context) {
	alunoID := c.Param("id")
	log.Printf("Recebendo requisição para editar aluno com ID: %s", alunoID)

	var input models.DetailsAluno
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verifica se o aluno existe
	var aluno models.Aluno
	if err := config.DB.First(&aluno, alunoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	// Cria o registro de detalhes do aluno e associa ao Aluno
	details := models.DetailsAluno{
		AlunoID:   aluno.ID,
		Cpf:       input.Cpf,
		Estado:    input.Estado,
		Cidade:    input.Cidade,
		Bairro:    input.Bairro,
		Rua:       input.Rua,
		Descricao: input.Descricao,
	}

	if err := config.DB.Create(&details).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, details)
}
