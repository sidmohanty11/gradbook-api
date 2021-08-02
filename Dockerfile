FROM golang:1.16

WORKDIR /usr/src/gradbook-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build
RUN ls
RUN cd ..
RUN ls

CMD ["server"]