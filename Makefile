.PHONY: binary clean shell coverage test

PACKAGES := $(shell go list ./... | grep -v /vendor/)

bin/twenty48: *.go cmd/twenty48/*.go
	@go build -o bin/twenty48 ./cmd/twenty48/...

binary: bin/twenty48

clean:
	@rm -rf ./bin coverage.txt || true

shell:
	@docker build -t twenty48:latest .
	@docker run -it -v $(shell pwd)/bin:/go/src/github.com/alyyousuf7/twenty48/bin twenty48:latest

coverage:
	@for pkg in $(PACKAGES); do \
		go test -race -coverprofile="../../../$$pkg/coverage.txt" -covermode=atomic $$pkg || exit; \
	done

test:
	@go test -v ./...
