PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test
test: lint
	go test $(PKGS)

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

$(GOMETALINTER):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.44.0

dev-up:
	docker-compose -f ./docker-compose.yml up -d

dev-recreate:
	docker-compose -f ./docker-compose.yml up --force-recreate --build -d

dev-down:
	docker-compose -f ./docker-compose.yml down

dev-clean: dev-down
	docker-compose -f ./docker-compose.yml rm -f

.PHONY: lint
lint: $(GOMETALINTER)
	golangci-lint run

BINARY := connectRN
VERSION ?= vlatest
PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64 ./cmd/server/main.go

.PHONY: release
release: windows linux darwin

.PHONY: clean
clean:
	rm -rf release/