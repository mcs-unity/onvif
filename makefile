# command flags
flags=-race -v

# build
output=bin/main.out

run:
	go run $(flags) cmd/main.go

test:
	go test ./...

build:test
	go build $(flags) -o $(output) cmd/main.go

