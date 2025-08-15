# Step 1: Build the binary
FROM golang:1.22 AS builder
WORKDIR /app

# Cache modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Step 2: Create small runtime image
FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=builder /app/server .
EXPOSE 8080
USER nonroot:nonroot

CMD ["./server"]