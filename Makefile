in ?= testdata
out ?= gen

run:
	go run ./cmd/hackasm -d $(in) -o $(out)

build:
	go build -o hackasm ./cmd/hackasm

test:
	go test ./...