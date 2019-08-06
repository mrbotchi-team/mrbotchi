FROM golang:1.12 AS dev

WORKDIR /go/src/github.com/mrbotchi-team/mrbotchi
VOLUME /go/src/github.com/mrbotchi-team/mrbotchi

RUN go get -u github.com/oxequa/realize

CMD [ "realize", "start", "--run" ]

FROM golang:1.12 AS build

WORKDIR /go/src/github.com/mrbotchi-team/mrbotchi

RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock makefile ./
RUN make deps

COPY . .
RUN make

FROM alpine:latest AS prod

EXPOSE 3000

WORKDIR /bin
VOLUME /.config

COPY --from=build /go/src/github.com/mrbotchi-team/mrbotchi/bin/mrbotchi .

CMD [ "./mrbotchi" ]
