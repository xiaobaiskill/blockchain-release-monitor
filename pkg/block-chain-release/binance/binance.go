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
)

type bsc struct {
	version   string
	sleepTime time.Duration
}

func NewBsc(sleepTime time.Duration) *bsc {
	return &bsc{
		sleepTime: sleepTime,
	}
}

func (e *bsc) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		e.getVersionToChan(bc)
		tc := time.NewTicker(e.sleepTime)
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
