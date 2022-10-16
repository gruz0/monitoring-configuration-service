.DEFAULT_GOAL := help

DOCKER_COMPOSE := docker compose

export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=monitoring-configuration-service
export LDFLAGS="-w -s"
export MONITORING_CONFIGURATION_DB_URL='host=localhost user=app password=password dbname=app_development sslmode=disable TimeZone=UTC'

help: # Show this help
	@egrep -h '\s#\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build:
	go build -race -o $(APP) .

build-static:
	CGO_ENABLED=0 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run: # Start web server
	go run -race . -db.url=$(MONITORING_CONFIGURATION_DB_URL)

test: # Run tests
	go clean -testcache
	MONITORING_CONFIGURATION_DB_URL='postgres://app:password@localhost:5433/app_test?TimeZone=UTC&sslmode=disable' \
	go test -v -race ./...

dockerize: # Run dockerized database and web server
	@${DOCKER_COMPOSE} up app db

docker-start-database:
	@${DOCKER_COMPOSE} up db

test-start-database: # Run dockerized test database
	@${DOCKER_COMPOSE} up test_db

docker-build: # Build container
	docker build --rm -t gruz0/monitoring-configuration-service .

.PHONY: build run build-static test dockerize docker-start-database test-start-database docker-build
