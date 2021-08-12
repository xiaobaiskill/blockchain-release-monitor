module github.com/xiaobaiskill/blockchain-release-monitor

go 1.15

replace github.com/xiaobaiskill/blockchain-release-monitor/api v0.0.0 => ./api

require (
	github.com/antchfx/htmlquery v1.2.3
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	github.com/xiaobaiskill/blockchain-release-monitor/api v0.0.0
	github.com/xiaobaiskill/kit v0.0.0-20210811141850-3d9434535cbd
	go.uber.org/zap v1.17.0
	google.golang.org/grpc v1.39.1
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)
