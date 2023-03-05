package test

import (
	"api/app/config"
	"api/app/helpers/requests"
	"api/app/repository"
	"strings"
	"testing"
)

var alunoUrl string

func TestCreateValidAluno(t *testing.T) {

	email := `"TesteGolangService@mail.com"`
	Db := config.DatabaseConnection()

	alunoRepository := repository.NewAlunosRepository(Db)
	aluno := alunoRepository.FindFirstBy(map[string]interface{}{
		"email": "TesteGolangService@mail.com",
	})

	if aluno != nil {
		alunoRepository.DeleteAluno(aluno)
	}

	requestBody := strings.NewReader(`
	{
	"name":"TesteGolangService",
	"nickname": "TesteGolangService",
	"email":` + email + `,
	"password":"TesteGolangService",
	"cellphone":"1112223333"
	}`)
	alunoUrl := "http://localhost:3000/alunos"
	res, body := requests.JsonPostRequest(alunoUrl, requestBody)

	if res.StatusCode == 400 {
		t.Error("Status 400, erro de validação com usuário valido")
	}

	if res.StatusCode != 201 {
		t.Error("Erro ao tentar criar Aluno", body)
	}

}

func TestCreateInvalidAluno(t *testing.T) {
	requestBody := strings.NewReader(`
	{
	"nickname": "testeInvalid",
	"email": "testeInvalid",
	"password":"testeInvalid",
	"cellphone":"1112223333"
	}`)
	alunoUrl := "http://localhost:3000/alunos"
	res, _ := requests.JsonPostRequest(alunoUrl, requestBody)

	if res.StatusCode != 400 {
		t.Error("Criouse aluno quando não deveria")
	}

}
