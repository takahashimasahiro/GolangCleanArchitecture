FROM golang:latest

WORKDIR /go/cmd
ADD . /go

CMD ["go", "run", "main.go"]

RUN go get github.com/go-sql-driver/mysql;