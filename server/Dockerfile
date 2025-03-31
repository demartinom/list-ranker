# Use the official Go image
FROM golang:1.21 AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source files and build
COPY . ./
WORKDIR /app/cmd
RUN go build -o /run-app .

# Use a lightweight runtime image
FROM gcr.io/distroless/base-debian11
COPY --from=builder /run-app /run-app

CMD ["/run-app"]
