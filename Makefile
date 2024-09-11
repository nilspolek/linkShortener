BINARY_NAME=bin/linkeShortener
GO_CMD = go

build:
	$(GO_CMD) build -o ./$(BINARY_NAME)

test:
	$(GO_CMD) test ./db/ . -count=1

clean:
	rm -f $(BINARY_NAME)

rmdb:
	rm -f ./link.db

run: build
	./$(BINARY_NAME)

lint:
	golint ./...

all: build

.PHONY: build test clean run lint fmt all
