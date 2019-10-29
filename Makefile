all: test

.PHONY: all

dep:
	go mod download

.PHONY: dep

build: dep
	go build -o bin/shamir cli/main.go

.PHONY: build

test: build
	./tests/shamir

.PHONY: test
