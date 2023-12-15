FROM debian:stable-slim AS final

LABEL version="1.0" \
      maintainer="carlosservi@correo.ugr.es"

# Actualiza la lista de paquetes e instala las dependencias necesarias
RUN apt-get update && apt-get install -y wget tar

# Descarga e instala Go
ENV GO_VERSION 1.21
RUN wget https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz && \
    rm go$GO_VERSION.linux-amd64.tar.gz

# Configura las variables de entorno
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin

WORKDIR /app

RUN adduser --disabled-password -u 1001 gouser

USER gouser

COPY go.mod go.sum ./

#Instalamos go mod y task
RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest

WORKDIR /app/test

ENTRYPOINT ["task", "test"]