all: build

build:
	go build -o server main.go

clean:
	rm server

run:
	go run main.go

test:
	export DEBUG=1 && go test -v

.PHONY: clean run
