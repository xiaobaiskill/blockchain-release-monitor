package service

import (
	"context"
	"github.com/xiaobaiskill/blockchain-release-monitor/internal/model"
	blockchain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	_ "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release/register"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/notification"
	"sync"
	"time"
)

type BlockChainer interface {
	GetReleaseData(ctx context.Context) []*model.ReleaseInfo
	ReleaseWatch(ctx context.Context)
}

var _ BlockChainer = (*blockChainImpl)(nil)

type blockChainImpl struct {
	bc             chan blockchain.BlockChainReleaseMsg
	releaseVersion map[string]model.ReleaseInfo
	sync.RWMutex
}

func NewblockChain() *blockChainImpl {
	return &blockChainImpl{
		bc:             make(chan blockchain.BlockChainReleaseMsg),
		releaseVersion: make(map[string]model.ReleaseInfo),
	}
}

func (b *blockChainImpl) GetReleaseData(_ context.Context) []*model.ReleaseInfo {
	b.RLock()
	out := make([]*model.ReleaseInfo, 0, len(b.releaseVersion))
	for _, v := range b.releaseVersion {
		info := v
		out = append(out, &info)
	}
	b.RUnlock()
	return out
}

func (b *blockChainImpl) ReleaseWatch(ctx context.Context) {
	blockchain.BlockChainMapRun(b.bc)
	go func() {
		for bc := range b.bc {
			if bc.Version == "" {
				continue
			}
			blockchainName := bc.Name + "-" + bc.Project
			releaseInfo, ok := b.releaseVersion[blockchainName]
			if ok {
				// compare updates, send notifications
				if releaseInfo.Version != bc.Version {
					b.Lock()
					b.releaseVersion[blockchainName] = model.ReleaseInfo{
						Name:    bc.Name,
						Project: bc.Project,
						Version: bc.Version,
						Url:     bc.Url,
						Time:    time.Now(),
					}
					b.Unlock()

					notification.RecvMsg(ctx, notification.Msg{
						Title: bc.Name,
						Level: 0,
						Data: map[string]interface{}{
							"project": bc.Project,
							"version": bc.Version,
							"url":     bc.Url,
						},
					})
				}
			} else {
				// For the first time, store directly
				b.Lock()
				b.releaseVersion[blockchainName] = model.ReleaseInfo{
					Name:    bc.Name,
					Project: bc.Project,
					Version: bc.Version,
					Url:     bc.Url,
					Time:    time.Now(),
				}
				b.Unlock()
			}
		}
	}()
}
