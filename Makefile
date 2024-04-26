build:
	rm -rf build
	GOARCH=amd64 GOOS=linux go build -o build/linux/xenv main.go
	GOARCH=amd64 GOOS=windows go build -o build/windows/xenv.exe main.go
	