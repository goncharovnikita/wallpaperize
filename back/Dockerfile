FROM golang:alpine as backdist

ENV dist=/go/src/github.com/goncharovnikita/wallpaperize/back

ADD . ${dist}

WORKDIR ${dist}

RUN apk add curl git --no-cache

RUN apk add --update gcc musl-dev

RUN env CGO_ENABLED=1 GOOS=linux GOARCH=386 go build --tags "libsqlite3 linux" -o ./back

FROM alpine

COPY --from=backdist /go/src/github.com/goncharovnikita/wallpaperize/back/back /app/app

WORKDIR /app

CMD [ "./app" ]
