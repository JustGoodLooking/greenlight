# 讀取環境（預設為 dev）
ENV ?= dev
include .env.$(ENV)
export


# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: confirm-deploy
confirm-deploy:
	@echo "⚠️  DEPLOYING with settings:"
	@echo "  - ENV        = $(ENV)"
	@echo "  - HOST       = $(HOST)"
	@echo "  - REMOTE_DIR = $(REMOTE_DIR)"
	@echo "  - VERSION    = $(VERSION)"
	@echo "  - DB_DSN     = $(GREENLIGHT_DB_DSN)"
	@echo -n "Are you sure you want to deploy? [y/N] " && read ans && [ $${ans:-N} = y ]



# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	docker-compose -f docker-compose.dev.yml up

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${KEEPLESS_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${KEEPLESS_DB_DSN} up

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format all .go files, and tidy and vendor module dependencies
.PHONY: tidy
tidy:
	@echo 'Tidying module dependencies...'
	go mod tidy
	@echo 'Verifying and vendoring module dependencies...'
	go mod verify
	go mod vendor
	@echo 'Formatting .go files...'
	go fmt ./...

## audit: run quality control checks
.PHONY: audit
audit:
	@echo 'Checking module dependencies...'
	go mod tidy -diff
	go mod verify
	@echo 'Vetting code...'
	go vet ./...
	go tool staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...



# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags="-s" -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o=./bin/linux_amd64/api ./cmd/api


## docker/build: build Docker image for api

VERSION := $(shell git describe --tags --always)
COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u +"%Y-%m-%d%H:%M:%SZ")

DOCKER_IMAGE_NAME := keepless-api

.PHONY: docker/build
docker/build:
	@echo "Building Docker image $(DOCKER_IMAGE_NAME):$(VERSION)"
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg BUILD_TIME=$(BUILD_TIME)\
		-t $(DOCKER_IMAGE_NAME):$(VERSION)



# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #


deploy:confirm-deploy docker-build upload migrate restart

upload:
	docker save $(DOCKER_IMAGE_NAME):$(VERSION) | bzip2 | ssh $(HOST) 'bunzip2 | docker load'
	rsync -avP .env.api.$(ENV) $(HOST):$(REMOTE_DIR)/.env.api.$(ENV)
	rsync -avP ./migrations/ $(HOST):$(REMOTE_DIR)/migrations/

migrate:
	ssh $(HOST) "\
		docker run --rm \
			-v $(REMOTE_DIR)/migrations:/migrations \
			migrate/migrate \
			-path=/migrations \
			-database '$(GREENLIGHT_DB_DSN)' \
			up \
	"

restart:
	ssh $(HOST) "\
		cd $(REMOTE_DIR) && \
		docker compose -f docker-compose.yml -f $(COMPOSE_FILE) down && \
		docker compose -f docker-compose.yml -f $(COMPOSE_FILE) up -d \
	"