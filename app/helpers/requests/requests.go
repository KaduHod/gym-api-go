package requests

import (
	"api/app/helpers/errors"
	"api/app/models"
	"encoding/json"
	"io"
)

type modelsGenrics interface {
	map[string]string | models.User | models.Personal | models.Aluno
}

func GetBodyJson[T modelsGenrics](body io.ReadCloser, target *T) {
	jsonBody, err := io.ReadAll(body)
	if err != nil {
		errors.Check(err)
	}
	json.Unmarshal([]byte(jsonBody), target)
}
