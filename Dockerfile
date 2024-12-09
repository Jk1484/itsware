FROM golang:1.22.0-alpine as build

RUN mkdir /itsware

ADD . /itsware

WORKDIR /itsware

RUN go build -o itsware ./cmd

FROM alpine:latest
COPY --from=build /itsware /itsware

WORKDIR /itsware

CMD ["/itsware/itsware"]