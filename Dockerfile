FROM alpine:latest

RUN apk add --no-cache tini chromium

COPY ./bin/ambot /ambot

ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/ambot"]
