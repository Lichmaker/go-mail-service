FROM golang:1.16.5 as goBuilder

RUN mkdir -p /mail-server
WORKDIR /mail-server
COPY . .
RUN /bin/bash -c 'if [ ! -e ".env" ]; then  echo "env file not found" ;  exit 1  ; else echo "success" ; exit 0; fi '
RUN go build -o mail-server

FROM alpine:latest
RUN mkdir -p /mail-server
WORKDIR /mail-server
COPY --from=goBuilder /mail-server/mail-server .
EXPOSE 3001
ENTRYPOINT ["./mail-server"]
