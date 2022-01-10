package cmd

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func (r *Request) PostRequest() (string, error) {
	response, err := http.Post(r.url, "application/json", bytes.NewBuffer([]byte(r.requestBody)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	byteArray, _ := ioutil.ReadAll(response.Body)
	return string(byteArray), nil
}
