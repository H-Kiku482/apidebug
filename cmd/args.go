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
	AutoExtFile         bool
}

func GetArgs() (*cmdArgs, error) {
	ca := new(cmdArgs)

	var (
		url                 string
		help                *bool
		method              *string
		requestBodyFilePath *string
		outputFile          *string
		autoExtFile         *bool
	)

	help = flag.Bool("h", false, "Print this message and exit.")
	method = flag.String("m", "GET", "Select http method.")
	requestBodyFilePath = flag.String("i", "", "Import request body from textfile.")
	outputFile = flag.String("o", "", "Output response body to textfile.")
	autoExtFile = flag.Bool("f", false, "Automaticary choose file extension. (json or html)")

	flag.Parse()

	url = flag.Arg(0)
	if !*help {
		if len(flag.Args()) == 0 {
			return nil, errors.New("not enough argument")
		} else if len(flag.Args()) != 1 {
			return nil, errors.New("too many arguments")
		}
	}

	ca.Help = *help
	ca.Url = url
	ca.Method = strings.ToUpper(*method)
	ca.RequestBodyFilePath = *requestBodyFilePath
	ca.OutputFilePath = *outputFile
	ca.AutoExtFile = *autoExtFile

	return ca, nil
}
