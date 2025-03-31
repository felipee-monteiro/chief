GOBUILDFLAGS = -ldflags "-s -w" -trimpath

.PHONY: build
build:
	rm -rf bin
	go build $(GOBUILDFLAGS) -o bin/ ./

.PHONY: run
run:
	go run .
