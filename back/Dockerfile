FROM golang:1.17 as backdist

ENV dist=/go/src/github.com/goncharovnikita/wallpaperize/back

ADD . ${dist}

WORKDIR ${dist}

RUN env CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -o ./back

FROM alpine

COPY --from=backdist /go/src/github.com/goncharovnikita/wallpaperize/back/back /app/app

WORKDIR /app

CMD [ "./app" ]
