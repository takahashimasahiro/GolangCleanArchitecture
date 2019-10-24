FROM golang:latest

WORKDIR /go/cmd
ADD . /go

CMD ["go", "run", "main.go"]