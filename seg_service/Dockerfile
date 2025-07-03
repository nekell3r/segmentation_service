# Build stage
FROM golang:1.24 AS builder
WORKDIR /app
COPY seg_service/go.mod seg_service/go.sum ./
RUN go mod download
COPY seg_service/. .
RUN go build -o seg_service main.go

# Final image
FROM golang:1.24
WORKDIR /app
COPY --from=builder /app/seg_service .
COPY ../.env .env
EXPOSE 8080
CMD ["./seg_service"]