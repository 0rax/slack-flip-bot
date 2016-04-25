FROM golang:1.6-alpine
MAINTAINER 0rax <jp@roemer.im>

RUN apk --no-cache --no-progress add git

COPY . /go/src/github.com/0rax/flip-bot
WORKDIR /go/src/github.com/0rax/flip-bot
RUN go get && go build

EXPOSE 4242
CMD ["/go/src/github.com/0rax/flip-bot/flip-bot"]
