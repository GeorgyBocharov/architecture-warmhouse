FROM golang:1.25-alpine AS builder

RUN apk add --no-cache make

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make generate

RUN CGO_ENABLED=0 GOOS=linux go build -o /application cmd/*

FROM alpine:latest

COPY --from=builder /application /application

ENTRYPOINT ["/application"]
