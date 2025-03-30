GOBUILDFLAGS = -ldflags "-s -w" -compiler gc -gcflags "-N -l"

.PHONY: build
build:
	rm -rf bin
	go build $(GOBUILDFLAGS) -o bin/ ./

.PHONY: run
run:
	go run .