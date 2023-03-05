package test

import (
	"api/app/helpers/requests"
	"api/app/models"
	"testing"
)

func TestGetOneAlunoService(t *testing.T) {
	res, body := requests.GetRequestWithJsonBody("http://localhost:3000/alunos?id=101")

	if res.StatusCode != 200 {
		t.Error("Erro ao fazer requisicao de aluno", res.StatusCode)
	}

	users := models.JsonToUsers[models.Aluno](body)

	if users[0].Id != 101 {
		t.Error("Pegou o use errado")
	}

}

func TestGetOnePersonalService(t *testing.T) {
	res, body := requests.GetRequestWithJsonBody("http://localhost:3000/personais?id=1")

	if res.StatusCode != 200 {
		t.Error("Erro ao fazer requisicao de personal", res.StatusCode)
	}

	users := models.JsonToUsers[models.Personal](body)

	if users[0].Id != 1 {
		t.Error("Pegou o use errado")
	}
}
