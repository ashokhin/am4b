GOVERSION	:= $(shell go env GOVERSION)
GOARCH		:= $(shell go env GOARCH)
GOOS		:= $(shell go env GOOS)

BIN_DIR		?= $(shell pwd)/bin/
BIN_NAME	?= $(shell go env GOEXE)

EXECUTABLE := ambot$(BIN_NAME)

export APP_HOST			?= $(shell hostname)
export APP_BRANCH		?= $(shell git describe --all --contains --dirty HEAD)
export APP_REVISION		?= $(shell git rev-parse HEAD)
export APP_ORIGIN		?= $(shell git config --local --get remote.origin.url)
export APP_VERSION		:= $(shell basename ${APP_BRANCH})
export APP_USER			:= $(shell id -u --name)
export APP_BUILD_DATE	:= $(shell date -u '+%Y-%m-%dT%H:%M:%S,%N%:z')

all: clean format vet test build

clean:
	@echo ">> removing build artifacts"
	@rm -f $(BIN_DIR)$(EXECUTABLE)

format:
	@echo ">> formatting code"
	@go fmt ./...

vet:
	@echo ">> vetting code"
	@go vet ./...

test:
	@echo ">> testing code"
	@go test ./...

linux: BIN_DIR=""
linux: clean format vet build

windows: BIN_DIR=""
windows: clean format vet build

build:
	@echo ">> building binary"
	@CGO_ENABLED=0 go build -v \
		-ldflags "-X github.com/prometheus/common/version.Version=${APP_VERSION} \
			-X github.com/prometheus/common/version.Branch=${APP_BRANCH} \
			-X github.com/prometheus/common/version.Revision=${APP_REVISION} \
			-X github.com/prometheus/common/version.BuildUser=${APP_USER}@${APP_HOST} \
			-X github.com/prometheus/common/version.BuildDate=${APP_BUILD_DATE} \
		" \
		-o ${BIN_DIR}${EXECUTABLE} ./cmd/ambot