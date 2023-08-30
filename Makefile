
# go params
GOCMD=go
GOBUILD=$(GOCMD) build 
GOTEST=$(GOCMD) test -run
GOPATH=/usr/local/bin
DIR=$(shell pwd)

build:
	@clear
	@echo "building pdf..."
	@$(GOBUILD) .

update:
	clear
	@echo "updating dependencies..."
	@$(GOCMD) get -u -t .
	@$(GOCMD) mod tidy 

test:
	@clear 
	@echo "testing QA functions..."
	@$(GOTEST) QA ./...

test-net:
	@clear 
	@echo "testing network functions..."
	@$(GOTEST) Net ./...
