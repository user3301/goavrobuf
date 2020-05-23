# Go parameters
    GOCMD=go
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get

.PHONY: lint
lint:
	golangci-lint -v run --timeout=1m

.PHONY: vendor
vendor:
	$(GOCMD) mod tidy && $(GOCMD) mod vendor

.PHONY: test
test:
	$(GOTEST) -cover ./...