all: ./cmd/main.go
	go build -o dist/dns-server ./cmd/main.go
clean:
	rm -rf dist/*