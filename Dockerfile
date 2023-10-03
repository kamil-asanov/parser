FROM golang:latest
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY src/ ./
RUN go run main.go
