
quickstart: src/*.go
	go build src/*.go
	
run:
	./src/quickstart

clean: 
	rm src/quickstart
