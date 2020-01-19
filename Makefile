default: build

build:
	go build -o dist/pssh cmd/main.go

install: build
	sudo cp dist/pssh /usr/local/bin/
