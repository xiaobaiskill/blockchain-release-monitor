package binance

import (
	blockchain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	githubRelease "github.com/xiaobaiskill/blockchain-release-monitor/pkg/github-release"
	"time"
)

const (
	name        = "binance smart chain"
	projectName = "binance-chain"
	url         = "https://github.com/binance-chain/bsc/releases/latest"
	sleepTime   = time.Hour
)

type bsc struct {
	version string
}

func init() {
	blockchain.RegisterBlockChainRelease(name, newBsc())
}

func newBsc() *bsc {
	return &bsc{}
}

func (e *bsc) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		e.getVersionToChan(bc)
		tc := time.NewTicker(sleepTime)
		for range tc.C {
			e.getVersionToChan(bc)
		}
	}()
}

func (e *bsc) getVersionToChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	version, _ := githubRelease.GetCurrentVersion(url)
	if version != "" {
		if e.version != version {
			e.version = version
			bc <- blockchain.BlockChainReleaseMsg{
				Name:    name,
				Project: projectName,
				Version: version,
				Url:     url,
			}
		}
	}
}
