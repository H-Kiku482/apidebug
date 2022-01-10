package cmd

func GetHelpText() string {
	return `Usage:  apidebug [OPTIONS]... URL
   or:  apidebug URL
   or:  apidebug -m POST -i FILE -o FILE URL

Options:
  -h    Print this message and exit.
  -i string
        Import request body from textfile.
  -m string
        Select http method. (default "GET")
  -o string
        Output response body to textfile.`
}
