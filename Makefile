.PHONY: all
all: build
FORCE: ;

SHELL := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-api

build-api:
	go build -tags $(LIBRARY_ENV) -o ./bin/api api/main.go

ci: dependencies test

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=usecase/template/interface.go -destination=usecase/template/mock/template.go -package=mock

test:
	go test -tags testing ./...

run-api:
	./bin/api

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done