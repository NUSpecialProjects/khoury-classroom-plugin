ARG GO_VERSION=1.23
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

ENTRYPOINT ["air"]
