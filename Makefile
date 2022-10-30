.PHONY: all test clean run

all: 
	mkdir -p ./dist/
	go build -o ./dist/ ./... 

test:
	go test ./...

test-clear:
	go clean -testcache

clean:
	rm wordle-bot
