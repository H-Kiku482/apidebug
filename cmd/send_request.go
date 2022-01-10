package cmd

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func (r *Request) SendRequest(method string) (string, error) {
	var (
		request *http.Request
		err     error
	)
	method = strings.ToUpper(method)

	if r.requestBody == "" {
		request, err = http.NewRequest(method, r.url, nil)
		if err != nil {
			return "", errors.New("failed to create a request")
		}
	} else {
		request, err = http.NewRequest(method, r.url, bytes.NewBuffer([]byte(r.requestBody)))
		if err != nil {
			return "", errors.New("failed to create a request")
		}
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", errors.New("failed to send a request")
	}
	defer response.Body.Close()

	byteArray, _ := ioutil.ReadAll(response.Body)
	return string(byteArray), nil
}
