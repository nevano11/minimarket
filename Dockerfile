FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -C /app/cmd/minimarket

EXPOSE 8080

ENTRYPOINT go run cmd/minimarket/main.go --config=deployment/local/minimarket.yaml