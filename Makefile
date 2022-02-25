lambda:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GIN_MODE=release go build -ldflags="-w -s" -o ./build/main main.go && cd build && zip -o ./lambda.zip ./main

deploy:
	aws lambda update-function-code --function-name alert-to-slack --zip-file fileb://build/lambda.zip --publish

all: lambda deploy

.PHONY: .lambda .deploy
