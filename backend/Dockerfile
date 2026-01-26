# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum* ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 creates a static binary compatible with distroless
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o resume-api .

# Production stage using distroless
# Uses static-debian12 for static Go binaries (no libc needed)
FROM gcr.io/distroless/static-debian12

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/resume-api .

# Distroless runs as non-root by default (uid 65532)
# Expose port
EXPOSE 8080

# Run the application
CMD ["./resume-api"]
