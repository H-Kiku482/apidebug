package cmd

import (
	"errors"
	"flag"
	"strings"
)

type cmdArgs struct {
	Help                bool
	Url                 string
	Method              string
	RequestBodyFilePath string
	OutputFilePath      string
}

func GetArgs() (*cmdArgs, error) {
	ca := new(cmdArgs)

	var (
		url                 string
		help                *bool
		method              *string
		requestBodyFilePath *string
		outputFile          *string
	)

	help = flag.Bool("h", false, "Print this message and exit.")
	method = flag.String("m", "get", "Select http method.")
	requestBodyFilePath = flag.String("i", "", "Import request body from textfile.")
	outputFile = flag.String("o", "", "Output response body to textfile.")

	flag.Parse()

	url = flag.Arg(0)
	if !*help {
		if len(flag.Args()) == 0 {
			return nil, errors.New("not enough argument")
		} else if len(flag.Args()) != 1 {
			return nil, errors.New("too many arguments")
		}
	}

	m := strings.ToLower(*method)
	method = &m

	switch *method {
	case "get":
	case "post":
	case "patch":
	case "delete":
	default:
		return nil, errors.New("cannot use the selected method\nplease select one of 'get', 'post', 'patch' or 'delete'")
	}

	ca.Help = *help
	ca.Url = url
	ca.Method = *method
	ca.RequestBodyFilePath = *requestBodyFilePath
	ca.OutputFilePath = *outputFile

	return ca, nil
}
