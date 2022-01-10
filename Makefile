native:
	go build

linux:
	GOOS=linux GOARCH=amd64 go build -o apidebug

windows:
	GOOS=windows GOARCH=amd64 go build -o apidebug.exe

clean:
	rm -rf apidebug apidebug.exe
