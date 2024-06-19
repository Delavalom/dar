.PHONY: run build

build: 
	go build -o dar cmd/dar/main.go 
	rm -rf ../dar_test/*
	mkdir -p ../dar_test/.dar
	cp -R ./.dar/* ../dar_test/.dar/*
	go build -o ../dar_test/dar cmd/dar/main.go

run:
	go run cmd/dar/main.go