# syntax=docker/dockerfile:1

## Build
FROM golang:1.20-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD ./cmd ./cmd
ADD ./internal ./internal

RUN GOARCH=amd64 GOOS=linux go build -o /getir ./cmd

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /getir /getir

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/getir"]