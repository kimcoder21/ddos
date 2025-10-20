# NamDoS Pro v2.0 Dockerfile
# Multi-stage build for optimal image size

# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY namdos_pro.go ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o namdos_pro namdos_pro.go

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN adduser -D -s /bin/sh namdos

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/namdos_pro .

# Change ownership to non-root user
RUN chown namdos:namdos namdos_pro

# Switch to non-root user
USER namdos

# Expose port (if needed for web interface)
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ./namdos_pro -test -site https://httpbin.org/get || exit 1

# Default command
CMD ["./namdos_pro"]
