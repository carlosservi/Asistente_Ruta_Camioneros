FROM golang:latest AS builder
FROM debian:stable-slim AS final

LABEL version="1.0" \
      maintainer="carlosservi@correo.ugr.es"

# Instalamos Go
COPY --from=golang:latest /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

# Incluimos los certificados
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN adduser --disabled-password -u 1001 gouser

USER gouser

WORKDIR /app/test

COPY go.mod go.sum ./

#Instalamos go mod y task
RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT ["task", "test"]