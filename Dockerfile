#
# 1. Build Container
#
FROM golang:1.15.6 AS build

ENV GO111MODULE=on \
    GOOS=linux \
    CGO_ENABLED=0 \
    GOARCH=amd64

RUN mkdir -p /src

# First add modules list to better utilize caching
COPY go.sum go.mod /src/

WORKDIR /src

# Download dependencies
RUN go mod download

COPY . /src

# Build components.
# Put built binaries and runtime resources in /app dir ready to be copied over or used.
RUN go install -installsuffix cgo -ldflags="-w -s" && \
    mkdir -p /app && \
    cp -r "${GOPATH}/bin/monitoring-configuration-service" /app/

#
# 2. Runtime Container
#
FROM alpine:3.13

LABEL maintainer="Alexander Kadyrov <gruz0.mail@gmail.com>"

ENV TZ=UTC \
    PATH="/app:${PATH}"

RUN apk add --update --no-cache \
    tzdata=2020f-r0 \
    ca-certificates=20191127-r5 \
    bash=5.1.0-r0 \
    && \
    cp --remove-destination /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone

# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# Create an user for running the application
RUN adduser -D app
USER app
WORKDIR /home/app

COPY --chown=app --from=build /app /home/app

EXPOSE 8080

CMD ["./monitoring-configuration-service"]
