FROM golang as build

WORKDIR /build

COPY . .

RUN go build -o myhive .

FROM alpine:latest
RUN apk add --no-cache \
        libc6-compat
COPY --from=build /build/myhive /usr/local/bin/myhive

EXPOSE 9000

ENTRYPOINT ["/usr/local/bin/myhive"]
