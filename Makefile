build:
	@go build -o bin/tv

run: build
	@./bin/tv

test:
	@go test ./... -v