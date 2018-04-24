all: build

.PHONY: all

build:
	go build -o bin/shamir cli/main.go

.PHONY: build

test:
	./tests/shamir

.PHONY: test
