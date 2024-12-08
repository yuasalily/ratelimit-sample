# ベースイメージ
FROM golang:1.23

# 作業ディレクトリを設定
WORKDIR /app

# モジュールファイルをコピーして依存関係を取得
COPY mock/* ./
RUN go mod download

# アプリケーションのビルド
RUN go build -o main .

# コンテナ実行時のコマンド
CMD ["./main"]