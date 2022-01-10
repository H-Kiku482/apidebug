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

	fmt.Println("HELP   :", args.Help)
	fmt.Println("URL    :", args.Url)
	fmt.Println("METHOD :", args.Method)
	fmt.Println("REQ    :", args.RequestBodyFilePath)
	fmt.Println("RES    :", args.OutputFilePath)

	request, err := cmd.NewRequest(args.Url, args.RequestBodyFilePath, args.OutputFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := ""
	switch m := args.Method; m {
	case "get":
		res, err = request.GetRequest()
		if err != nil {
			fmt.Println(err)
		}
	case "post":
		res, err = request.PostRequest()
		if err != nil {
			fmt.Println(err)
		}
	case "patch":
	case "delete":
	default:
		fmt.Println("internal error")
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
