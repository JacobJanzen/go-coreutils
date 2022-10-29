.PHONY: all test clean run

all: 
	mkdir -p ./dist/
	go build -o ./dist/ ./... 

test:
	go test ./...

clean:
	rm wordle-bot
