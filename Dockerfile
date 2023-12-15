FROM golang:latest AS builder
FROM debian:stable-slim AS final

LABEL version="1.0" \
      maintainer="carlosservi@correo.ugr.es"

COPY --from=golang:latest /usr/local/go/ /usr/local/go/

ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest

RUN adduser --disabled-password -u 1001 gouser

USER gouser

WORKDIR /app/test

ENTRYPOINT ["task", "test"]