BINARY_NAME=appbuilder-grader
OUTPUT_DIR=bin

all: build

build:
	go build -o ${OUTPUT_DIR}/${BINARY_NAME} main.go

build-all: build build-linux build-mac build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${OUTPUT_DIR}/${BINARY_NAME}-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o ${OUTPUT_DIR}/${BINARY_NAME}-linux-arm64 main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o ${OUTPUT_DIR}/${BINARY_NAME}-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o ${OUTPUT_DIR}/${BINARY_NAME}-darwin-arm64 main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ${OUTPUT_DIR}/${BINARY_NAME}-windows-amd64.exe main.go

clean:
	go clean
	rm -rf ${OUTPUT_DIR}

run: build
	./${OUTPUT_DIR}/${BINARY_NAME}
