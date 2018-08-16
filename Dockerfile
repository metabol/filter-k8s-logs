FROM alpine:3.7

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY bin/filter /usr/bin/filter
EXPOSE 8080

CMD /usr/bin/filter