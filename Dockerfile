# Use the golang image
FROM golang:1.22.3-alpine AS builder

# Set the working directory
WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o web-server src/main.go

FROM node:22.2 AS tailwind-builder
WORKDIR /app
COPY --from=builder /app .

FROM alpine:latest
WORKDIR /app

COPY --from=tailwind-builder /app/web-server .
COPY --from=tailwind-builder /app/static ./static
COPY --from=tailwind-builder /app/*.pem .
COPY --from=tailwind-builder /app/templates ./templates
COPY --from=tailwind-builder /app/config.yaml .

# Expose port 8080 for HTTP
EXPOSE 8080
# Expose port 8443 for HTTPS
EXPOSE 8443
CMD ["./web-server"]
