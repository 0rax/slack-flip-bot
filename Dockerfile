FROM alpine:3.3
MAINTAINER 0rax <jp@roemer.im>

COPY . /app
WORKDIR /app

ENV GOPATH /go
RUN apk --no-cache --no-progress add --virtual build-deps go git \
 && mkdir -p /go/src/github.com/0rax/ \
 && ln -s /app /go/src/github.com/0rax/slack-flip-bot \
 && go get github.com/0rax/slack-flip-bot && go build -o /app/flip-bot \
 && rm -rf /go/* && apk --no-cache --no-progress del build-deps

EXPOSE 4242
CMD ["/app/flip-bot"]
