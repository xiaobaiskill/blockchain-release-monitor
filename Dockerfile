FROM golang:1.14-alpine as builder
LABEL builder=blockchain-release-monitor

WORKDIR /src

ADD . .

RUN CGO_ENABLE=0 go build  -ldflags="-s -w" -race -o ./build/app .


FROM alpine
WORKDIR /
COPY --from=builder /src/build/ /
ENTRYPOINT ["/app"]
CMD ["server",".config.json"]