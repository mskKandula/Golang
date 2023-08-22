FROM golang:1.18-alpine as builder

# Install git
RUN apk update && \
    apk upgrade && \
    apk add git

# Set the Current Working Directory inside the container
WORKDIR /app/V3/server

# We want to populate the module cache based on the go.{mod,sum} files.
COPY ./server/go.mod .

COPY ./server/go.sum .

RUN go mod download

COPY ./server/ .

# Build the Go app
RUN go build -o main .

#Define the running image
FROM alpine:3.17

RUN apk update && \
    apk upgrade

# Set the Current Working Directory inside the container
WORKDIR /app/V3/server

# Copy built binary application from 'builder' image
COPY --from=builder /app/V3/server/main .

# This container exposes port 8000 to the outside world
EXPOSE 8000

# Run the binary program produced by `go install`
CMD ["./main"]