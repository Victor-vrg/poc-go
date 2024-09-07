FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /poc-go

# Stage final
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /poc-go .
EXPOSE 3000
CMD ["./poc-go"]
