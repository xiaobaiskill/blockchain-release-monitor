package heco

import (
	blockchain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	githubRelease "github.com/xiaobaiskill/blockchain-release-monitor/pkg/github-release"
	"time"
)

const (
	name        = "heco"
	projectName = "huobi-eco-chain"
	url         = "https://github.com/HuobiGroup/huobi-eco-chain/releases/latest"
)

type heco struct {
	version   string
	sleepTime time.Duration
}

func NewHeco(sleepTime time.Duration) *heco {
	return &heco{sleepTime: sleepTime}
}

func (e *heco) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		e.getVersionToChan(bc)
		tc := time.NewTicker(e.sleepTime)
		for range tc.C {
			e.getVersionToChan(bc)
		}
	}()
}

func (e *heco) getVersionToChan(bc chan<- blockchain.BlockChainReleaseMsg) {
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
