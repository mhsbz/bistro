FROM golang:alpine

MAINTAINER dxj

WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

COPY . .

RUN go mod tidy && go build cmd/main.go

CMD ["./main"]

