package test

import (
	"api/app/helpers/requests"
	"api/app/models"
	"testing"
)

func TestGetExerciseByIdWithExistingId(t *testing.T) {
	res, body := requests.GetRequestWithJsonBody("http://localhost:3000/exercicios/1330")

	if res.StatusCode != 200 {
		t.Error("Erro ao buscar exercicio por id")
	}

	exercises := models.JsonToExercise(body)

	if exercises.Id != 1330 {
		t.Error("Trouxe o exercicio errado")
	}

}

func TestGetExerciseByIdWithNonExistingId(t *testing.T) {
	res := requests.GetRequest("http://localhost:3000/exercicios/99999")

	if res.StatusCode != 404 {
		t.Error("BUscou exercicio quando nao deveria")
	}
}
