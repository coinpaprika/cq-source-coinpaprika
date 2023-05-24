.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables

.PHONY: gen-mocks
gen-mocks:
	# go install github.com/golang/mock/mockgen
	rm -rf ./client/mocks/*
	go generate ./client/...

# All gen targets
.PHONY: gen
gen: gen-docs gen-mocks