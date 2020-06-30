FROM golang:latest

WORKDIR /go/src/todo/api/application

RUN go get -u github.com/labstack/echo/...
RUN go get github.com/go-xorm/xorm  
RUN go get github.com/go-sql-driver/mysql

EXPOSE 8880