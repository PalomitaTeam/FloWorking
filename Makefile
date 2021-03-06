SRC   = src
CXX   = go
FLAGS = build

compile:
	go build src/*.go
	
run:
	go run src/*.go

clean: 
	go clean