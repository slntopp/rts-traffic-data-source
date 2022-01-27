FROM golang:1.17

COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod download

ENTRYPOINT [ "go", "run", "/app/main.go" ]