########################################################
# ビルドステージ
########################################################
FROM golang:1.24-bookworm AS builder

WORKDIR /app

# 依存関係ファイルをコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

########################################################
# 実行ステージ
########################################################
FROM debian:bullseye-slim

WORKDIR /app

# ビルドステージから実行ファイルをコピー
COPY --from=builder /app/main .

# アプリケーションの実行
EXPOSE 8080
CMD ["./main"] 
