all: gomod vet test

gomod:
	go mod tidy
	go mod verify

vet:
	go vet ./...

test:
	go test ./... -timeout 30s -race