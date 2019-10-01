FROM golang:latest

WORKDIR /go/src/po_users
COPY . .

RUN go get -d ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["po_users"]
