package register

import (
	block_chain_release "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release/avalanche"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release/binance"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release/etherum"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release/etherum2.0"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release/heco"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release/okexchain"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release/polygon"
	"time"
)

const (
	sleepTime = time.Minute * 20
)

func init() {
	block_chain_release.RegisterBlockChainRelease(
		avalanche.NewAvalanche(sleepTime),
		binance.NewBsc(sleepTime),
		etherum.NewEth(sleepTime),
		etherum2.NewEth2(sleepTime),
		heco.NewHeco(sleepTime),
		okexchain.NewOkex(sleepTime),
		polygon.NewPolygon(sleepTime),
		polygon.NewPolygon(sleepTime),
	)
}
