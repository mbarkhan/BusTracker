# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.19.0-alpine3.16 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build ./auth/main.go
##
## Deploy
##
# FROM gcr.io/distroless/base-debian10
FROM alpine:latest 
WORKDIR /

COPY --from=build ./app/main ./main

ENTRYPOINT ["./main"]