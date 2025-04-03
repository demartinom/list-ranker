# Use the official Go image as the builder
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go.mod and go.sum to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . ./

# Copy the game-data folder to the app directory
COPY ./game-data /app/game-data/

# Ensure a statically linked binary
RUN CGO_ENABLED=0 go build -o /run-app ./cmd

# Use a minimal runtime image
FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=builder /run-app /run-app

# Ensure the game-data folder is copied
COPY --from=builder /app/game-data /app/game-data

# Use a non-root user for security
USER nonroot

# Run the application
CMD ["/run-app"]
