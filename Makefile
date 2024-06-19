UNAME := $(shell uname -o)


# The /dar_test dir is for testing purposes
build:
	go build -o dar cmd/dar/main.go 
    ifeq ($(UNAME), Darwin)
		if [ -d "$(HOME)/bin" ]; then \
			cp ./dar $(HOME)/bin/; \
    	fi
    endif

	rm -rf ../dar_test/*
	mkdir -p ../dar_test/.dar
	cp -R ./.dar/* ../dar_test/.dar/*

run:
	go run cmd/dar/main.go

.PHONY: run build