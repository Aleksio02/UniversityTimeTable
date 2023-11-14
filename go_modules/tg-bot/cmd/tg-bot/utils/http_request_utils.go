package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func CreatePathParam(paramName string, paramValue string) string {
	return paramName + "=" + paramValue
}

func WriteBodyToObject(body io.Reader, objectToWrite any) {
	responseBody, _ := ioutil.ReadAll(body)
	json.Unmarshal(responseBody, &objectToWrite)
}
