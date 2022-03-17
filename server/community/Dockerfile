FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build main.go

EXPOSE 8080

CMD ["./main"]

