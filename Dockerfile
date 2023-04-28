FROM  golang:1.18.1-alpine3.14

COPY . /app

WORKDIR /app

RUN go build -o main main.go

#CMD ["/main"]

