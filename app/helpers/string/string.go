package string

import (
	"encoding/json"
	"strconv"
)

func ToInteger(value string) (int, error) {
	integer, err := strconv.Atoi(value)
	return integer, err
}

func JsonToMapOfArrays[T any](jsonBody string) map[string][]T {
	var result map[string][]T
	json.Unmarshal([]byte(jsonBody), &result)
	return result
}

func JsonToMap(jsonBody string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonBody), &result)
	return result
}
