FROM golang:alpine as backdist

ENV dist=/go/src/github.com/goncharovnikita/wallpaperize/back

ADD . ${dist}

WORKDIR ${dist}

RUN apk add curl git clang clang-dev musl-dev git gcc --no-cache

ENV CXX=clang++

ENV CC=clang

RUN env CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -o ./back

FROM alpine

COPY --from=backdist /go/src/github.com/goncharovnikita/wallpaperize/back/back /app/app

WORKDIR /app

CMD [ "./app" ]
