package test

import (
	"api/app/helpers/requests"
	"strings"
	"testing"
)

func TestUpdateValidAlunoService(t *testing.T) {
	email := "TesteGolangAlunoUpdateService@mail.com"
	url := "http://localhost:3000/alunos"
	requestBody := strings.NewReader(`
	{
	"id":102,
	"name":"TesteGolangAlunoUpdateService2",
	"nickname": "TesteGolangAlunoUpdateService",
	"email": "` + email + `",
	"password":"TesteGolangAlunoUpdateService",
	"cellphone": 1112223333
	}`)

	res, body := requests.JsonPutRequest(url, requestBody)

	if res.StatusCode != 200 {
		t.Error("Erro ao atualizar aluno", body)
	}
}

func TestUpdateValidPersonalService(t *testing.T) {
	email := "TesteGolangPersonalUpdateService@mail.com"
	url := "http://localhost:3000/personais"
	requestBody := strings.NewReader(`
	{
	"id":1,
	"name":"TesteGolangPersonalUpdateService",
	"nickname": "TestGolangPersonalUpdService",
	"email": "` + email + `",
	"password":"TesteGolangPersonalUpdateService",
	"cellphone": 1112223333
	}`)

	res, body := requests.JsonPutRequest(url, requestBody)

	if res.StatusCode != 200 {
		t.Error("Erro ao atualizar Personal", body)
	}
}

func TestUpdateInvalidPersonalService(t *testing.T) {
	email := "TesteGolangPersonalUpdateService@mail.com"
	url := "http://localhost:3000/personais"
	requestBody := strings.NewReader(`
	{
	"name":"TesteGolangPersonalUpdateService2",
	"nickname": "TestGolangPersonalUpdService",
	"email": "` + email + `",
	"password":"TesteGolangPersonalUpdateService",
	"cellphone": 1112223333
	}`)

	res, body := requests.JsonPutRequest(url, requestBody)

	if res.StatusCode != 400 {
		t.Error("Erro ao atualizar Personal quando não deveria", body)
	}
}

func TestUpdateInvalidAlunoService(t *testing.T) {
	email := "TesteGolangAlunoUpdateService@mail.com"
	url := "http://localhost:3000/alunos"
	requestBody := strings.NewReader(`
	{
	"name":"TesteGolangAlunoUpdateService",
	"nickname": "TesteGolangAlunoUpdateService",
	"email": "` + email + `",
	"password":"TesteGolangAlunoUpdateService",
	"cellphone": 1112223333
	}`)

	res, body := requests.JsonPutRequest(url, requestBody)

	if res.StatusCode != 400 {
		t.Error("Erro ao atualizar aluno quando não deveria", body)
	}
}
