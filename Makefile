CLI_BINARY_NAME=appbuilder-grader
LAMBDA_BINARY_NAME=appbuilder-grader-lambda
OUTPUT_DIR=bin

all: build-cli build-lambda

build: build-cli

build-cli:
	go build -o ${OUTPUT_DIR}/${CLI_BINARY_NAME} ./cmd/cli

build-lambda:
	GOOS=linux GOARCH=amd64 go build -o ${OUTPUT_DIR}/${LAMBDA_BINARY_NAME}-linux-amd64 ./cmd/lambda
	GOOS=linux GOARCH=arm64 go build -o ${OUTPUT_DIR}/${LAMBDA_BINARY_NAME}-linux-arm64 ./cmd/lambda

build-all: build-cli build-lambda build-linux build-mac build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${OUTPUT_DIR}/${CLI_BINARY_NAME}-linux-amd64 ./cmd/cli
	GOOS=linux GOARCH=arm64 go build -o ${OUTPUT_DIR}/${CLI_BINARY_NAME}-linux-arm64 ./cmd/cli

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o ${OUTPUT_DIR}/${CLI_BINARY_NAME}-darwin-amd64 ./cmd/cli
	GOOS=darwin GOARCH=arm64 go build -o ${OUTPUT_DIR}/${CLI_BINARY_NAME}-darwin-arm64 ./cmd/cli

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ${OUTPUT_DIR}/${CLI_BINARY_NAME}-windows-amd64.exe ./cmd/cli

clean:
	go clean
	rm -rf ${OUTPUT_DIR}

run: build
	./${OUTPUT_DIR}/${CLI_BINARY_NAME}
