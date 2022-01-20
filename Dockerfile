# syntax=docker/dockerfile:1
# Start from golang base image
FROM golang:alpine as builder

ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Igor Ariza <igorariza@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:3.4
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.envrc .


# Expose port 8080 to the outside world
EXPOSE 8000


CMD ["./main"]

#syntax=docker/dockerfile:1
# FROM golang:alpine AS build
# ENV GOPROXY=https://proxy.golang.org
# RUN apk add --no-cache gcc g++ git openssh-client curl
# WORKDIR /api

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download 
# COPY . .

# RUN go build  ./cmd/beer
# EXPOSE 8080

# CMD ["./main"]