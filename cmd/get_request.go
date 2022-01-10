package cmd

import (
	"io/ioutil"
	"net/http"
)

func (r *Request) GetRequest() (string, error) {
	response, err := http.Get(r.url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	byteArray, _ := ioutil.ReadAll(response.Body)
	return string(byteArray), nil
}
