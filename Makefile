lint:
	golangci-lint run --sort-results --show-stats --allow-parallel-runners

run: lint
	go run cmd/main/main.go

build: lint
	go build -o cmd/bin/main.exe cmd/main/main.go