build-client:
	go build -o bin/client ./client

build-server:
	go build -o bin/server ./server

build: build-client build-server

.PHONY: build build-client build-server