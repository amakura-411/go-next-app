# ベースイメージを指定
FROM golang:1.22.2

# COPY . /backend/my-app

# コンテナ内に作業ディレクトリを作成
WORKDIR /backend/my-app

RUN go install github.com/cosmtrek/air@latest


# CMD [ "go", "run", "main.go"]

# ホストのファイルをコンテナの作業ディレクトリに移行
# COPY . .

# 永続
# CMD [ "ls" ]

# CMD [ "ls" ]
# ポート番号を環境変数に設定
# EXPOSE 8080

# RUN go mod init


# # goのモジュールをインストール

CMD ["air", "-c", ".air.toml"]
