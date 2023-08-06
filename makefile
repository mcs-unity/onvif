# go run flags
rflags=-race -v
# go build flags
bflags= -v
# build
output=bin/main
# repository
repo=onvif

install:
	go mod download && go mod verify

run:
	go run $(rflags) cmd/main.go

test:
	go test ./...

build:test
	go build $(bflags) -o $(output) cmd/main.go

build_docker:
	docker build --secret id=env,src=.env -f docker/dockerfile -t $(repo) .

run_docker:
	docker run $(repo)

