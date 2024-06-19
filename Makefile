UNAME := $(shell uname -o)


build:
	go build -o dar cmd/dar/main.go 
    ifeq ($(UNAME), Darwin)
		if [ -d "$(HOME)/bin" ]; then \
			cp ./dar $(HOME)/bin/; \
    	fi
    endif

run:
	go run cmd/dar/main.go

# The /dar_test dir is for testing purposes
clone:
	rm -rf ../dar_test/*
	if [ -d "../dar_test/.dar" ]; then \
		cp -R ./.dar/ ../dar_test/.dar/; \
	else \
		mkdir ../dar_test/.dar; \
		cp -R ./.dar/ ../dar_test/.dar/; \
	fi

.PHONY: run build clone