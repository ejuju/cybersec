all: gomod test

gomod:
	go mod tidy
	go mod verify

test:
	go vet
	go test ./... -cover -timeout 60s -race -cpu 4