# ===== 第一階段：build Go binary =====
FROM golang:1.22 as builder

WORKDIR /app
COPY . .

# 接收外部參數（來自 docker build --build-arg）
ARG VERSION
ARG COMMIT
ARG BUILD_TIME

RUN go build -ldflags "\
    -X 'main.version=${VERSION}' \
    -X 'main.commit=${COMMIT}' \
    -X 'main.buildTime=${BUILD_TIME}'" \
    -o server ./cmd/api

# ===== 第二階段：生成最小化映像 =====
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/server .

CMD ["./server"]
