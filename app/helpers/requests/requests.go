package requests

import (
	"api/app/helpers/errors"
	"api/app/models"
	"encoding/json"
	"io"
)

type ModelsGenerics interface {
	map[string]string | models.User | models.Personal | models.Aluno
}

func GetBodyJson[T models.UserTypes](body io.ReadCloser, target *T) {
	jsonBody, err := io.ReadAll(body)

	if err != nil {
		errors.Check(err)
	}
	json.Unmarshal([]byte(jsonBody), &target)
}
