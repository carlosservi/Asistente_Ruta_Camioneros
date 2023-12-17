# Etapa de construcción
FROM golang:latest AS builder

COPY --from=golang:latest /usr/local/go/ /usr/local/go/

ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest && adduser --disabled-password -u 1001 gouser

# Elimina lo que ya no necesito
RUN rm -rf /usr/local/go/src /usr/local/go/pkg /usr/local/go/test /usr/local/go/api /usr/local/go/doc

# Etapa final
FROM debian:stable-slim AS final

LABEL version="1.0" \
      maintainer="carlosservi@correo.ugr.es"

COPY --from=builder /usr/local/go/ /usr/local/go/ /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ /home/gouser /home/gouser/

USER gouser

WORKDIR /app/test

ENTRYPOINT ["task", "test"]
