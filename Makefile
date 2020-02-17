.PHONY: binary clean shell coverage test

PACKAGES := $(shell go list ./...)

bin/twenty48: *.go cmd/twenty48/*.go
	@go build -o bin/twenty48 ./cmd/twenty48/...

binary: bin/twenty48

clean:
	@rm -rf ./bin coverage.txt || true

shell:
	@docker build -t twenty48:latest .
	@docker run -it -v $(shell pwd)/bin:/twenty48/bin twenty48:latest

coverage:
	@for pkg in $(PACKAGES); do \
		go test -race -coverprofile=coverage.txt -covermode=atomic $$pkg || exit; \
	done

test:
	@go test -v ./...
