export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=monitoring-configuration-service
export LDFLAGS="-w -s"
export MONITORING_CONFIGURATION_DB_URL='host=localhost user=app password=password dbname=app_development sslmode=disable TimeZone=UTC'

all: build test

build:
	go build -race -o $(APP) .

build-static:
	CGO_ENABLED=0 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run -race . -db.url=$(MONITORING_CONFIGURATION_DB_URL)

test:
	go clean -testcache
	MONITORING_CONFIGURATION_DB_URL='postgres://app:password@localhost:5433/app_test?TimeZone=UTC&sslmode=disable' \
	go test -v -race ./...

dockerize:
	docker-compose up --build

build-container:
	docker build -t gruz0/monitoring-configuration-service .

run-container:
	docker run --rm -it gruz0/monitoring-configuration-service

.PHONY: build run build-static test dockerize build-container run-container
