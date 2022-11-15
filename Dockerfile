FROM golang:1.19-alpine3.16

COPY . .

ENV GOPATH=/

RUN go mod download

RUN go build -o balance-service ./cmd/api/main.go

CMD ["./balance-service"]
