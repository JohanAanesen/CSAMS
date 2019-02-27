#first stage - builder
FROM golang:1.11.0-stretch as builder

COPY . /WebService

WORKDIR /WebService

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o WebService

#second stage

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache tzdata

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /WebService .

EXPOSE 8088

CMD ["./WebService"]