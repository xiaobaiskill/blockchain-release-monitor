package service

import (
	"context"
)

var (
	BlockChainSvc BlockChainer
)

func GlobalInit() {
	BlockChainSvc = NewblockChain()

	BlockChainSvc.ReleaseWatch(context.Background())
}
