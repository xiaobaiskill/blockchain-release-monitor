FROM golang:1.14-alpine as builder
LABEL builder=blockchain-release-monitor

WORKDIR /src

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./build/app .

FROM alpine
WORKDIR /
COPY --from=builder /src/build/ /
ENTRYPOINT ["/app"]
CMD ["server",".config.json"]