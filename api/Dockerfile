# Use Go official image
FROM golang:1.25.1-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all files
COPY . .

# Expose API port
EXPOSE 2712

# Start dev server with hot reload
CMD ["air", "-c", ".air.toml"]

