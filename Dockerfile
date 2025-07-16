# 第一階段：編譯 Go 程式
FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go build -o server ./cmd/api

# 第二階段：建立最小化的執行環境（無 Go 環境）
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server"]