package polygon

import (
	blockchain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	githubRelease "github.com/xiaobaiskill/blockchain-release-monitor/pkg/github-release"
	"time"
)

const (
	name         = "polygon"
	borName      = "bor"
	borUrl       = "https://github.com/maticnetwork/bor/releases/latest"
	heimdallName = "heimdall"
	heimdallUrl  = "https://github.com/maticnetwork/heimdall/releases/latest"
)

type polygon struct {
	borVersion, heimdallVersion string
	sleepTime                   time.Duration
}

func NewPolygon(sleepTime time.Duration) *polygon {
	return &polygon{
		sleepTime: sleepTime,
	}
}

func (p *polygon) RunWithChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	go func() {
		p.getVersionOfBorToChan(bc)
		p.getVersionOfHeimdallToChan(bc)
		tc := time.NewTicker(p.sleepTime)
		for range tc.C {
			p.getVersionOfBorToChan(bc)
			p.getVersionOfHeimdallToChan(bc)
		}
	}()
}

func (p *polygon) getVersionOfBorToChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	version, _ := githubRelease.GetCurrentVersion(borUrl)
	if version != "" {
		if p.borVersion != version {
			p.borVersion = version
			bc <- blockchain.BlockChainReleaseMsg{
				Name:    name,
				Project: borName,
				Version: version,
				Url:     borUrl,
			}
		}
	}
}

func (p *polygon) getVersionOfHeimdallToChan(bc chan<- blockchain.BlockChainReleaseMsg) {
	version, _ := githubRelease.GetCurrentVersion(heimdallUrl)
	if version != "" {
		if p.heimdallVersion != version {
			p.heimdallVersion = version
			bc <- blockchain.BlockChainReleaseMsg{
				Name:    name,
				Project: heimdallName,
				Version: version,
				Url:     heimdallUrl,
			}
		}
	}
}
