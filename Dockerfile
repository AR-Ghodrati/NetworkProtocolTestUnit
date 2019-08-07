FROM golang:1.12.7-alpine3.10

RUN apk update && apk add git

WORKDIR /go/src/gsm
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

EXPOSE 3001

CMD ["gsm"]