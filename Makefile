.PHONY: build clean run

# Verify that the go.mod file is up to date, and then tidy it up.
# Build the binary for the playground app and place it in the bin directory.
build:
	go mod verify && go mod tidy
	go build -ldflags="-s -w" -o bin/playground ./playground/main.go

clean:
	rm -rf ./bin

# Runs the api lambda locally on whatever port is specified in the API_PORT environment variable.
# Concurrently runs the playground app on port 8080.
run: clean build
	./bin/playground &
	serverless offline start --httpPort ${API_PORT}
	