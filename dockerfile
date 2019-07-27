
FROM golang:1.12 AS build

WORKDIR /go/src/github.com/mrbotchi-team/mrbotchi

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN make deps && make

FROM golang:1.12 AS dev

WORKDIR /go/src/github.com/mrbotchi-team/mrbotchi
VOLUME /go/src/github.com/mrbotchi-team/mrbotchi

RUN go get -u github.com/golang/dep/cmd/dep github.com/oxequa/realize
CMD [ "realize", "start", "--run" ]

FROM alpine:latest AS prod

EXPOSE 3000

WORKDIR /bin

COPY --from=build /go/src/github.com/mrbotchi-team/mrbotchi/bin/mr-bochi-be .

CMD [ "./mr-bochi-be" ]
