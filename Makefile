all: get lint test

lint:
	go fmt ./...
get:
	go get -v -t -d ./...
test:
	go test ./...
deps:
	go list -m -u all
