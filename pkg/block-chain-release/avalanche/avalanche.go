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
	sleepTime   = time.Hour
)

type avalanche struct {
	version string
}

func init() {
	blockchain.RegisterBlockChainRelease(name, newAvalanche())
}

func newAvalanche() *avalanche {
	return &avalanche{}
}

func (e *avalanche) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		e.getVersionToChan(bc)
		tc := time.NewTicker(sleepTime)
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
