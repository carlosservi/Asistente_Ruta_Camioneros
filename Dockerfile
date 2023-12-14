FROM golang:1.21-alpine3.19

WORKDIR /app/test

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT ["task"]
CMD ["test"]