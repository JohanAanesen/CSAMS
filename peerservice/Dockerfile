#first stage - builder
FROM golang:1.11.0-stretch as builder

COPY . /PeerService

WORKDIR /PeerService

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o PeerService

#second stage

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache tzdata

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /PeerService .

EXPOSE 8087

CMD ["./PeerService"]