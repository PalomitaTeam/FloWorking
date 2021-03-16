
build: src/*.go
	go build src/*.go
	
run:
	./src/quickstart

check:
	go vet ./src/quickstart

test: 
	go test src/*_test.go

clean: 
	rm src/quickstart
