FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /FRECNET
RUN cd /FRECNET && make FRE

FROM alpine:latest

WORKDIR /FRECNET

COPY --from=builder /FRECNET/build/bin/FRE /usr/local/bin/FRE

RUN chmod +x /usr/local/bin/FRE

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/FRE"]

CMD ["--help"]
