include .env

## help: available commands
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run/app: run cli
.PHONY: run/app
run/app:
	@./bin/cli

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/app: production build
.PHONY: build/app
build/app:
	@echo 'Building cmd/app...'
	go build  -o=./bin/rss-feed-generator-cli ./cmd/cli


## audit: run quality control
.PHONY: audit
audit: vendor
	@echo "Formatting code..."
	go fmt ./...
	@echo "Vetting code..."
	go vet ./...
	staticcheck ./...
	@echo "Running tests..."
	go test -race -vet=off ./...

.PHONY: vendor
vendor:
	@echo "Tidying and verifying module dependecies..."
	go mod tidy
	go mod verify
	@echo "Vendoring dependecies..."
	go mod vendor
