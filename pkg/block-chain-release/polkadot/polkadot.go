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
	sleepTime   = time.Hour
)

type polkadot struct {
	version string
}

func init() {
	blockchain.RegisterBlockChainRelease(name, newPolkadot())
}

func newPolkadot() *polkadot {
	return &polkadot{}
}

func (e *polkadot) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		e.getVersionToChan(bc)
		tc := time.NewTicker(sleepTime)
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
