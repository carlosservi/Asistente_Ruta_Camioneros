FROM golang:1.21-alpine

RUN adduser --disabled-password -u 1001 user

USER user

WORKDIR /app/test

COPY go.mod go.sum ./

RUN go mod download && go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT ["task"]
CMD ["test"]