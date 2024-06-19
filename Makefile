UNAME := $(shell uname -o)


# The /dar_test dir is for testing purposes
build:
	go build -o dar cmd/dar/main.go 
    ifeq ($(UNAME), Darwin)
		if [ -d "$(HOME)/bin" ]; then \
			cp ./dar $(HOME)/bin/; \
    	fi
    endif

run:
	go run cmd/dar/main.go

clone:
	rm -rf ../dar_test/*
	mkdir ../dar_test/.dar
	cp -R ./.dar/ ../dar_test/.dar/

.PHONY: run build clone