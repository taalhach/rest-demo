all: build-rest

build-rest:
	mkdir -p bin
	go build -o bin/rest cmd/rest/main.go
	
