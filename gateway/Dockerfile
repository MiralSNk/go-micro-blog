FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /gateway ./cmd/gateway

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /gateway .
EXPOSE 8080 50051

CMD ["./gateway"]