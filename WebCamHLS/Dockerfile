FROM golang:1.21-alpine

# Install git
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ffmpeg

# Set the Current Working Directory inside the container
WORKDIR /app/DockerUsage

COPY . .

# Build the Go app
RUN go build -o main .


# This container exposes port 8080 to the outside world
EXPOSE 8081

# Run the binary program produced by `go install`
CMD ["./main"]