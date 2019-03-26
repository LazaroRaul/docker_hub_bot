FROM golang:1.10-stretch

WORKDIR /go/src/docker_hub_bot
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["cmd"]
