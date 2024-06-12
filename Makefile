.PHONY: run build

build: 
	go build -o dar cmd/dar/main.go

run:
	go run cmd/dar/main.go