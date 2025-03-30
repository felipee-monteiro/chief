GOBUILDFLAGS = -ldflags "-s -w" -compiler gc -gcflags "-N -l"

.PHONY: build
build:
	go build $(GOBUILDFLAGS) -o bin/ ./

.PHONY: run
run:
	go run .