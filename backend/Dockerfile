# ベースイメージを指定
FROM golang:1.22.2

# コンテナ内に作業ディレクトリを作成
WORKDIR /backend/my-app

# ポート番号を環境変数に設定
EXPOSE 1323

RUN go mod init

RUN go get -u github.com/cosmtrek/air

RUN go mod download

# # goのモジュールをインストール

CMD ["air", "-c", ".air.toml"]
