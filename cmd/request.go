package cmd

import (
	"errors"
	"io/ioutil"
	"os"
)

type Request struct {
	url              string
	requestBody      string
	responseFilePath string
}

func NewRequest(url string, requestBodyFilePath string, responseFilePath string) (*Request, error) {
	r := new(Request)

	if requestBodyFilePath != "" {
		f, err := os.Open(requestBodyFilePath)
		if err != nil {
			return nil, errors.New("failed to open " + requestBodyFilePath)
		}
		defer f.Close()

		data, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, errors.New("failed to open " + requestBodyFilePath)
		}

		r.requestBody = string(data)
	} else {
		r.requestBody = ""
	}

	r.url = url
	r.responseFilePath = responseFilePath

	return r, nil
}
