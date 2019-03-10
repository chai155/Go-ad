# Generic makefile. Developed and tested with the windows platform.

# Runs only a build for the local architecture in
# the local directory when "make build" is run

build:
	@go build

test:
	@go test