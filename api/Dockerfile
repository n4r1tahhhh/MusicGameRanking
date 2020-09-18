FROM golang:latest

RUN mkdir -p /go/src/api
WORKDIR /go/src/api
COPY . .

# 他にもモジュール管理やホットリロードなど
RUN go get -u github.com/labstack/echo/... && \ 
    go get -u gorm.io/gorm && \
    go get -u gorm.io/driver/mysql && \
    go get -u github.com/joho/godotenv && \
    go get -u github.com/dgrijalva/jwt-go

EXPOSE 1323