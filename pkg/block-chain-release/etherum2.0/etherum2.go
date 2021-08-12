package etherum2

import (
	blockchain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	githubRelease "github.com/xiaobaiskill/blockchain-release-monitor/pkg/github-release"
	"time"
)

const (
	name        = "etherum2.0"
	projectName = "prysm"
	url         = "https://github.com/prysmaticlabs/prysm/releases/latest"
	sleepTime   = time.Hour
)

type eth2 struct {
	version string
}

func init() {
	blockchain.RegisterBlockChainRelease(name, newEth2())
}

func newEth2() *eth2 {
	return &eth2{}
}

func (e *eth2) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		e.getVersionToChan(bc)
		tc := time.NewTicker(sleepTime)
		for range tc.C {
			e.getVersionToChan(bc)
		}
	}()
}

func (e *eth2) getVersionToChan(bc chan<- blockchain.BlockChainReleaseMsg) {
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
