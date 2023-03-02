package test

import (
	"api/app/config"
	"api/app/repository"
	"fmt"
	"testing"
)

var alunoUrl string

func TestCreateValidAluno(t *testing.T) {
	// alunoUrl := "http://localhost:3000/alunos"
	email := `"TesteGolangService@mail.com"`
	Db := config.DatabaseConnection()
	alunoRepository := repository.NewAlunosRepository(Db)
	aluno := alunoRepository.FindFirstBy(map[string]string{
		"email": email,
	})

	fmt.Println(aluno)

	// requestBody := strings.NewReader(`
	// {
	// "name":"TesteGolangService",
	// "nickname": "TesteGolangService",
	// "email":` + email + `,
	// "password":"TesteGolangService",
	// "cellphone":"1112223333"
	// }`)
	//
	// fmt.Println(requestBody, "request")
	// res, body := requests.JsonPostRequest(alunoUrl, requestBody)
	//
	// if res.Status == "400" {
	// t.Error("Status 400, erro de validação com usuário valido")
	// }
	//
	// if res.Status != "201" {
	// t.Error("Erro ao tentar criar Aluno", body)
	// }

}

func TestCreateInvalidAluno(t *testing.T) {

}
