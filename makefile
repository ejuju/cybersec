all: gomod vet test

gomod:
	@printf "\nTidying go modules: "
	go mod tidy
	
	@printf "\nVerifying go modules: "
	go mod verify

vet:
	@printf "\nVetting code: "
	go vet ./...

test:
	@printf "\nRunning tests: "
	go test ./... -timeout 30s -race