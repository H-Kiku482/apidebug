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

	if args.OutputFilePath == "" && args.AutoExtFilePath == "" {
		fmt.Println(res)
	} else {
		if args.AutoExtFilePath != "" {
			switch r := res[0:1]; r {
			case "{":
				args.AutoExtFilePath += ".json"
			default:
				args.AutoExtFilePath += ".html"
			}

			f, err := os.Create(args.AutoExtFilePath)
			if err != nil {
				fmt.Println("failed to output " + args.AutoExtFilePath)
			}

			f.WriteString(res)
			f.Close()
		}
		if args.OutputFilePath != "" {
			f, err := os.Create(args.OutputFilePath)
			if err != nil {
				fmt.Println("failed to output " + args.OutputFilePath)
			}
			f.WriteString(res)
			f.Close()
		}
	}
}
