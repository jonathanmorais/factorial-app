FROM golang:1.13 AS build-env

LABEL maintainer="Jonathan Morais <jonathan.m.lucena@gmail.com>"

WORKDIR /go/src/app/

#install dependecies
RUN go get github.com/gorilla/mux
    # go get github.com/mattn/go-sqlite3 && \
    # go get github.com/sirupsen/logrus

#copy to workdir path
COPY ./src/main.go .

ENV CGO_ENABLED=0

#build the go app
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

RUN chmod +x /go/src/app/main.go


# final stage
FROM alpine
RUN apk --no-cache add ca-certificates

WORKDIR /app/

#copy the compilate binary for workdir
COPY --from=build-env /go/src/app .

ENTRYPOINT ["./app"]