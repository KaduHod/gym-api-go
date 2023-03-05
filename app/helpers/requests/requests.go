package requests

import (
	"api/app/helpers/errors"
	"api/app/models"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func JsonToIoReader(stringJson []byte) *bytes.Reader {
	return bytes.NewReader(stringJson)
}

func SetStructToBody(structData any) *bytes.Buffer {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(structData)
	if err != nil {
		log.Fatal(err)
	}

	return &buf
}

func GetBodyFromResponse(res *http.Response) string {
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func GetBodyFromRequest(body io.ReadCloser) string {
	json, err := ioutil.ReadAll(body)

	if err != nil {
		panic(err)
	}

	return string(json)
}

func JsonPostRequest(url string, body *strings.Reader) (*http.Response, string) {
	client, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		panic(err)
	}
	client.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(client)

	if err != nil {
		panic(err)
	}

	return res, GetBodyFromRequest(res.Body)
}

func JsonPutRequest(url string, body *strings.Reader) (*http.Response, string) {
	client, err := http.NewRequest(http.MethodPut, url, body)

	if err != nil {
		panic(err)
	}
	client.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(client)

	if err != nil {
		panic(err)
	}

	return res, GetBodyFromRequest(res.Body)
}

func GetRequest(url string) *http.Response {
	client, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(client)
	if err != nil {
		panic(err)
	}
	return res

}
