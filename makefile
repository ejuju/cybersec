all: gomod test

gomod:
	go mod tidy
	go mod verify

test:
	go test ./... -cover -timeout 30s -race -cover -vet "" -cpu 4