package block_chain_release

var blockChainReleaseMap = make([]blockChainReleaser, 0, 10)

func RegisterBlockChainRelease(bc ...blockChainReleaser) {
	blockChainReleaseMap = append(blockChainReleaseMap, bc...)
}

func BlockChainMapRun(blockChainMsgCh chan<- BlockChainReleaseMsg) {
	for _, bc := range blockChainReleaseMap {
		bc.RunWithChan(blockChainMsgCh)
	}
}
