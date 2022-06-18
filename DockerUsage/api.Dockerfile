FROM golang:1.16-alpine

# Install git
RUN apk update && \
    apk upgrade && \
    apk add git

# Set the Current Working Directory inside the container
WORKDIR /app/DockerUsage

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./app/DockerUsage .


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./app/DockerUsage"]