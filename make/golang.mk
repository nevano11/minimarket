.PHONY: go/install-deps
go/install-deps:
	brew install golang-migrate
	brew install golangci-lint

.PHONY: go/lint
go/lint:
	export CGO_ENABLE=0 && \
	golangci-lint run -v

.PHONY: go/fix
go/fix:
	export CGO_ENABLE=0 && \
	golangci-lint run -v --fix

.PHONY: go/test
go/test:
	export CGO_ENABLE=0 && \
	go test ./...

.PHONY: go/race
go/race:
	export CGO_ENABLE=1 && \
	go test -race ./...