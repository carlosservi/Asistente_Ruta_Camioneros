FROM golang:1.21-alpine

LABEL version="1.0" \
      maintainer="carlosservi@correo.ugr.es"

RUN adduser --disabled-password -u 1001 gouser

USER gouser

WORKDIR /app/test

COPY go.mod go.sum ./

RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT ["task", "test"]