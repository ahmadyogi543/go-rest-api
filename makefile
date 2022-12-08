run:
	CompileDaemon -build="go build -o ./bin/go-rest-api" -command="./bin/go-rest-api"

build:
	go build -o ./bin/go-rest-api

clean:
	rm -rf ./database/*