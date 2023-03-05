package test

import (
	"api/app/config"
	"api/app/helpers/requests"
	"api/app/repository"
	"strings"
	"testing"
)

var personalUrl string

func TestCreateValidPersonal(t *testing.T) {

	email := `TesteGolangPersonalService@mail.com`
	Db := config.DatabaseConnection()

	personalRepository := repository.NewPersonalRepository(Db)
	personal := personalRepository.FindFirstBy(map[string]interface{}{
		"email": email,
	})

	if personal.Id != 0 {
		personalRepository.DeletePersonal(personal)
	}

	requestBody := strings.NewReader(`
	{
	"name":"TesteGolangPersonalService",
	"nickname": "TesteGolangPersonalService",
	"email": "TesteGolangPersonalService@mail.com",
	"password":"TesteGolangPersonalService",
	"cellphone":"1112223333"
	}`)
	personalUrl := "http://localhost:3000/personais"
	res, body := requests.JsonPostRequest(personalUrl, requestBody)

	if res.StatusCode == 400 {
		t.Error("Status 400, erro de validação com usuário valido")
		return
	}

	if res.StatusCode != 201 {
		t.Error("Erro ao tentar criar Personal", body)
	}

}

func TestCreateInvalidPersonal(t *testing.T) {
	requestBody := strings.NewReader(`
	{
	"nickname": "testeInvalid",
	"email": "testeInvalid",
	"password":"testeInvalid",
	"cellphone": 1112223333
	}`)
	alunoUrl := "http://localhost:3000/personais"
	res, _ := requests.JsonPostRequest(alunoUrl, requestBody)

	if res.StatusCode != 400 {
		t.Error("Criou-se personal quando não deveria")
	}

}
