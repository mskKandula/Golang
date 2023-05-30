# Stage 1 as builder
FROM golang:1.18-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app/SPH/server

# Copy Go Application Dependencies
COPY ./Server/go.mod .
COPY ./Server/go.sum .

# To keep only actual dependencies
RUN go mod tidy

# Download dependencies 
RUN go mod download

# Copy code files
COPY ./Server/ .

# Build the Go app
RUN go build -o main .

########### Stage 2 ###########

#Define the running image
FROM alpine:3.17

RUN apk update && \
    apk upgrade

# Set the Current Working Directory inside the container
WORKDIR /app/SPH/server

# Copy built binary application from 'builder' image
COPY --from=builder /app/SPH/server/main .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./main"]