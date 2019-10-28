# 開発環境
FROM golang:1.13 AS dev

WORKDIR /go/src/github.com/mrbotchi-team/mrbotchi
VOLUME /go/src/github.com/mrbotchi-team/mrbotchi

RUN go get -u github.com/oxequa/realize

CMD [ "realize", "start", "--run" ]

FROM golang:1.13 AS build

WORKDIR /go/src/github.com/mrbotchi-team/mrbotchi

COPY . .
RUN make

# 本番環境
FROM alpine:latest AS prod

EXPOSE 3000

WORKDIR /bin
VOLUME /.config

COPY --from=build /go/src/github.com/mrbotchi-team/mrbotchi/bin/mrbotchi .

CMD [ "./mrbotchi" ]

