#syntax=docker/dockerfile:1
FROM golang:alpine AS build
ENV GOPROXY=https://proxy.golang.org
RUN apk add --no-cache gcc g++ git openssh-client curl
WORKDIR /api

COPY go.mod ./
COPY go.sum ./

RUN go mod download 
COPY . .

RUN go build  ./cmd/beer
EXPOSE 8080

CMD ["./gogrpc-apibeer"]