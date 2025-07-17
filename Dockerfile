FROM golang:1.22 as builder

WORKDIR /app
COPY . .

RUN go build -o server ./cmd/api


FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/server .

CMD ["./server"]
