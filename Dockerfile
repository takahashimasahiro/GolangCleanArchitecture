FROM golang:latest

WORKDIR /go/cmd
ADD . /go

CMD ["go", "build", "main.go"]
CMD ["go", "run", "main.go"]

RUN go get github.com/go-sql-driver/mysql \
    && go get github.com/google/uuid \
    && go get github.com/joho/godotenv;
COPY .env .env

# ENV PATH $PATH:
# CMD ["cd", "db/"]
# CMD ["docker-compose", "build"]