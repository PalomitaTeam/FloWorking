
quickstart: src/*.go
	go build src/*.go
	
run:
	./src/quickstart

check:
	go vet src/

clean: 
	rm src/quickstart
