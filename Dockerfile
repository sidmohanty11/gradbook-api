FROM golang:1.16

WORKDIR /usr/src/gradbook-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go get github.com/pilu/fresh

CMD ["gradbook-api"]