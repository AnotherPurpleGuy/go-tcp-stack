# syntax=docker/dockerfile:1

#########################
# Build stage (builder)
#########################
FROM golang:1.24 AS builder
WORKDIR /app

# Copy go.mod and source code
COPY go.mod .
COPY main.go .

# Download dependencies and build the binary statically.
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o tcpstack .

#########################
# Final stage (runtime)
#########################
FROM scratch
COPY --from=builder /app/tcpstack /tcpstack

ENTRYPOINT ["/tcpstack"]
