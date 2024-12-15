# ベースイメージ
FROM golang:1.23.4-alpine3.21

# 作業ディレクトリを設定
WORKDIR /app

# モジュールファイルをコピーして依存関係を取得
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main .

# コンテナ実行時のコマンド
CMD ["./main"]
