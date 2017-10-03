FROM alpine:3.5

RUN apk add --update curl jq && rm -rf /var/cache/apk/*

COPY ./go-app /usr/local/bin
COPY ./hostIP.sh /usr/local/bin

ENTRYPOINT ["go-app"]