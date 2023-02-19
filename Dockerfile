# syntax=docker/dockerfile:1

# Build
FROM golang:alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /servicefinder cmd/servicefinder/main.go

# Deploy
FROM alpine

WORKDIR /

COPY --from=build /servicefinder /servicefinder


ENTRYPOINT [ "/servicefinder" ]
