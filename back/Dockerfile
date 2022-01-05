FROM golang:alpine as backdist

ENV dist=/go/src/github.com/goncharovnikita/wallpaperize/back

ADD . ${dist}

WORKDIR ${dist}

RUN apk add curl --no-cache && \
    apk add git --no-cache

RUN apk --no-cache add gcc libc-dev

RUN env CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -o ./back

FROM alpine

COPY --from=backdist /go/src/github.com/goncharovnikita/wallpaperize/back/back /app/app

WORKDIR /app

CMD [ "./app" ]
