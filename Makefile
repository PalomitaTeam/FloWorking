
quickstart: src/*.go
	go build src/*.go
	
run:
	./src/quickstart

check:
	go vet (or sth similar)

clean: 
	rm src/quickstart
