FROM golang:1.8-alpine AS builder

COPY . /go/src/github.com/0rax/slack-flip-bot

RUN cd /go/src/github.com/0rax/slack-flip-bot \
 && CGO_ENABLED=0 GOOS=linux go build -o bin/flip-bot -ldflags '-extldflags "-static"'

FROM scratch AS runtime

WORKDIR /app
COPY --from=builder /go/src/github.com/0rax/slack-flip-bot/bin/* /app/bin/

EXPOSE 4242
CMD ["bin/flip-bot"]
