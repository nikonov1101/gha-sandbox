FROM golang:1.16-alpine as builder

RUN apk add git make

WORKDIR /build
COPY . .
RUN make build/linux

FROM alpine:latest
COPY --from=builder /build/fooer /fooer
CMD ["/fooer"]
