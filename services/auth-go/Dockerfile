FROM golang:1.23.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o auth .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/auth .
EXPOSE 8080
CMD ["./auth"]
