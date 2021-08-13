package avalanche

import (
	blockchain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	githubRelease "github.com/xiaobaiskill/blockchain-release-monitor/pkg/github-release"
	"time"
)

const (
	name        = "avalanche"
	projectName = "avalanchego"
	url         = "https://github.com/ava-labs/avalanchego/releases/latest"
)

type avalanche struct {
	version   string
	sleepTime time.Duration
}

func NewAvalanche(sleepTime time.Duration) *avalanche {
	return &avalanche{
		sleepTime: sleepTime,
	}
}

func (e *avalanche) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		e.getVersionToChan(bc)
		tc := time.NewTicker(e.sleepTime)
		for range tc.C {
			e.getVersionToChan(bc)
		}
	}()
}

func (e *avalanche) getVersionToChan(bc chan<- blockchain.BlockChainReleaseMsg) {
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
