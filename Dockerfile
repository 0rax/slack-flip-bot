FROM golang:1.6-alpine
MAINTAINER 0rax <jp@roemer.im>

RUN apk --no-cache --no-progress add git

COPY . /go/src/github.com/0rax/slack-flip-bot
WORKDIR /go/src/github.com/0rax/slack-flip-bot
RUN go get && go build -o flip-bot

EXPOSE 4242
CMD ["/go/src/github.com/0rax/slack-flip-bot/flip-bot"]
