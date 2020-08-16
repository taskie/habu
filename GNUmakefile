.PHONY: build

build:
	go build -v -ldflags "-s -w"

.PHONY: install

install:
	go install -v -ldflags "-s -w"

.PHONY: rebuild-examples

rebuild-examples:
	$(MAKE) -C examples/features rebuild
	$(MAKE) -C examples/subcommands rebuild
