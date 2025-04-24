# --- Build Stage ---
    FROM golang:1.21-bullseye AS builder

    WORKDIR /app
    
    # Cache dependencies
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the rest of the application
    COPY . .
    
    # Build the Go application
    RUN go build -v -o server ./cmd/main.go
    
    # --- Final Stage ---
    FROM debian:bullseye-slim
    
    # Optional: Install certificates (required if your app makes HTTPS calls)
    RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && \
        rm -rf /var/lib/apt/lists/*
    
    # Copy the binary and necessary files
    COPY --from=builder /app/server /app/server
    COPY --from=builder /app/config/config-docker.yml /app/config/config-docker.yml
    COPY --from=builder /app/docs /app/docs
    
    # Set environment variables
    ENV APP_ENV=docker
    ENV PORT=8080
    
    # Run the binary
    CMD ["/app/server"]
    