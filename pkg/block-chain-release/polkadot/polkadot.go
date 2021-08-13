package polkadot

import (
	blockchain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	githubRelease "github.com/xiaobaiskill/blockchain-release-monitor/pkg/github-release"
	"time"
)

const (
	name        = "polkadot"
	projectName = "polkadot"
	url         = "https://github.com/paritytech/polkadot/releases/latest"
)

type polkadot struct {
	version   string
	sleepTime time.Duration
}

func NewPolkadot(sleepTime time.Duration) *polkadot {
	return &polkadot{sleepTime: sleepTime}
}

func (e *polkadot) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		e.getVersionToChan(bc)
		tc := time.NewTicker(e.sleepTime)
		for range tc.C {
			e.getVersionToChan(bc)
		}
	}()
}

func (e *polkadot) getVersionToChan(bc chan<- blockchain.BlockChainReleaseMsg) {
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
