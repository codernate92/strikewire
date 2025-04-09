#Dockerfile
FROM golang:1.21 as builder

WORKDIR /app
COPY . . 
RUN go build -o strikewire ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/strikewire .
ENTRYPOINT ["./strikewire"]