FROM golang:alpine as backdist

ENV dist=/go/src/github.com/goncharovnikita/wallpaperize/back

ADD . ${dist}

WORKDIR ${dist}

RUN apk add curl --no-cache && \
    apk add git --no-cache

RUN env GOOS=linux GOARCH=386 go build

FROM alpine

COPY --from=backdist /go/src/github.com/goncharovnikita/wallpaperize/back/back /app/app

WORKDIR /app

CMD [ "./app" ]
