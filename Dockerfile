FROM golang:1.9

WORKDIR /go/src/github.com/wlwanpan/web-services
ADD . .

RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2

RUN go build server.go

EXPOSE 80

CMD ["./server"]