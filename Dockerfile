FROM golang:1.16.5-alpine3.13 as goBuilder
RUN apk add --no-cache bash
RUN mkdir -p /mail-service
WORKDIR /mail-service
COPY . .
RUN /bin/bash -c 'if [ ! -e ".env" ]; then  echo "env file not found" ;  exit 1  ; else echo "success" ; exit 0; fi '
RUN go build -o mail-service

FROM alpine:latest
RUN mkdir -p /mail-service
WORKDIR /mail-service
COPY --from=goBuilder /mail-service/mail-service .
COPY --from=goBuilder /mail-service/.env .
ENTRYPOINT ["./mail-service"]
