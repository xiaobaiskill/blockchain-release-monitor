package okexchain

import (
	blockchain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	githubRelease "github.com/xiaobaiskill/blockchain-release-monitor/pkg/github-release"
	"time"
)

const (
	name        = "okexchain"
	projectName = "exchain"
	url         = "https://github.com/okex/exchain/releases/latest"
	sleepTime   = time.Hour
)

type okex struct {
	version string
}

func init() {
	blockchain.RegisterBlockChainRelease(name, newOkex())
}

func newOkex() *okex {
	return &okex{}
}

func (e *okex) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		e.getVersionToChan(bc)
		tc := time.NewTicker(sleepTime)
		for range tc.C {
			e.getVersionToChan(bc)
		}
	}()
}

func (e *okex) getVersionToChan(bc chan<- blockchain.BlockChainReleaseMsg) {
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
