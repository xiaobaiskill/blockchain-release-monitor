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
)

type okex struct {
	version   string
	sleepTime time.Duration
}

func NewOkex(sleepTime time.Duration) *okex {
	return &okex{sleepTime: sleepTime}
}

func (e *okex) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		e.getVersionToChan(bc)
		tc := time.NewTicker(e.sleepTime)
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
