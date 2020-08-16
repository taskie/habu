.PHONY: build

build:
	go build -v -ldflags "-s -w"

.PHONY: generate

generate:
	go generate

.PHONY: rebuild

rebuild: generate build
