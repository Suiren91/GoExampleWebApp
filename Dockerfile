FROM golang:1.26.5-bookworm as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# -----------------------------------------------------

# デプロイ用コンテナ
FROM debian:bookworm-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# -----------------------------------------------------

# ホットリロード環境(ローカル)
FROM golang:1.26.5 as dev

WORKDIR /app

RUN go install github.com/air-verse/air@latest
CMD ["air"]
