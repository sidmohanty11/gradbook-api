FROM golang:1.16.6-alpine

WORKDIR /usr/src/gradbook-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

CMD ["server"]