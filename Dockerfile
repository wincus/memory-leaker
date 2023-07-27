# Start from the latest Golang base image
FROM golang:1.17-alpine AS builder

# Add Maintainer Info
LABEL maintainer="Juan Moyano <wincus@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod main.go ./

# Build the Go app
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Command to run the executable
ENTRYPOINT ["./main"]