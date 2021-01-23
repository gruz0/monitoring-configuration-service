export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=monitoring-configuration-service
export LDFLAGS="-w -s"

all: build test

build:
	go build -race -o $(APP) .

build-static:
	CGO_ENABLED=0 go build -race -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run -race .

test:
	go test -v -race ./...

dockerize:
	docker-compose up --build

build-container:
	docker build -t gruz0/monitoring-configuration-service .

run-container:
	docker run --rm -it gruz0/monitoring-configuration-service

.PHONY: build run build-static test dockerize build-container run-container
