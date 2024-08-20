package models

import "gorm.io/gorm"

// entidade de aluno
type Aluno struct {
	gorm.Model
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
}
type DetailsAluno struct {
	gorm.Model
	AlunoID   uint   `json:"aluno_id"` // Chave estrangeira para o Aluno
	Cpf       string `json:"cpf"`
	Estado    string `json:"estado"`
	Cidade    string `json:"cidade"`
	Bairro    string `json:"bairro"`
	Rua       string `json:"rua"`
	Descricao string `json:"descricao"`
}
