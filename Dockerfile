FROM golang:1.14-alpine as builder
LABEL builder=blockchain-release-monitor

WORKDIR /src
RUN go env -w GOPROXY="https://goproxy.cn,direct"

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./build/blockchain-release-monitor .

FROM alpine
WORKDIR /
COPY --from=builder /src/build/ /
COPY --from=builder /src/config.json /config.json
ENTRYPOINT ["/blockchain-release-monitor"]
CMD ["server","/config.json"]