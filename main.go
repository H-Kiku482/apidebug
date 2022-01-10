package main

import (
	"apidebug/cmd"
	"fmt"
	"os"
)

func main() {
	args, err := cmd.GetArgs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if args.Help {
		fmt.Println(cmd.GetHelpText())
		os.Exit(0)
	}

	request, err := cmd.NewRequest(args.Url, args.RequestBodyFilePath, args.OutputFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := request.SendRequest(args.Method)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if args.OutputFilePath == "" {
		fmt.Println(res)
	} else {
		f, err := os.Create(args.OutputFilePath)
		if err != nil {
			fmt.Println("failed to output " + args.OutputFilePath)
		}
		f.WriteString(res)
		f.Close()
	}
}
