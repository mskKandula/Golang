FROM golang:1.18-alpine as builder

# Install git
RUN apk update && \
    apk upgrade && \
    apk add git

# Set the Current Working Directory inside the container
WORKDIR /app/V3/client

# We want to populate the module cache based on the go.{mod,sum} files.
COPY ./client/go.mod .

COPY ./client/go.sum .

RUN go mod download

COPY ./client .

# Build the Go app
RUN go build -o main .

#Define the running image
FROM alpine:3.17

# echo 10000 65535 > /proc/sys/net/ipv4/ip_local_port_range

# WORKDIR /etc/sysctl.d

# Copy configuration
# COPY ../portExtend/net.ipv4.ip_local_port_range.conf /etc/sysctl.d/

RUN apk update && \
    apk upgrade

# Set the Current Working Directory inside the container
WORKDIR /app/V3/client

# Copy built binary application from 'builder' image
COPY --from=builder /app/V3/client/main .

# Run the binary program produced by `go install`
CMD ["./main"]