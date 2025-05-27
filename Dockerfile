FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

# Run tests
RUN go test ./...

# Build (if you have a main package)
# RUN go build -o main ./cmd/main.go

# Keep the final image as golang instead of alpine to have Go available
FROM golang:1.21-alpine

WORKDIR /app

# Copy the package files
COPY --from=builder /app/pkg ./pkg
COPY --from=builder /app/go.mod .

# Default command (modify as needed)
CMD ["sh"] 