# Build stage
FROM node:alpine AS node-builder

# Install tailwindcss
RUN npm install -g tailwindcss

FROM golang:1.24.2-alpine AS builder

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@v0.3.865

# Install build dependencies
RUN apk add --no-cache make nodejs npm

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Install tailwindcss
RUN npm install -g tailwindcss@3.4.17

# Copy source code
COPY . .

# Run the build
RUN make build

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy required assets
COPY --from=builder /app/main .
COPY --from=builder /app/cmd/web/assets/css/output.css ./cmd/web/assets/css/output.css

# Expose the port your app runs on
EXPOSE 8080

ENV GOTRACEBACK=all

# Run the binary
CMD ["./main"]
