dist:
	./scripts/dist.sh

run:
	go run cmd/tmt/main.go

fmt:
	gofmt -w internal cmd
