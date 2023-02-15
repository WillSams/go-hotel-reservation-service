.PHONY: build clean run

# Ehhh...got to used to how handy npm scripts are in NodeJS, so I'm using a Makefile here.  
# This is seemingly a thing for Go-nistas (ewwww...Go-ninjas? Go-nauts? Go-nauts sounds cool)
# so I'm going with it.

# Nothing to really build/clean here, but I'm leaving this in for now.
# Let's at least make sure we have the dependencies installed.
build:
	go mod tidy

run: build
	serverless offline start --httpPort 8080