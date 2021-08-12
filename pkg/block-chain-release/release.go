package block_chain_release

var blockChainReleaseMap = make(map[string]blockChainReleaser)

func RegisterBlockChainRelease(name string, bc blockChainReleaser) {
	blockChainReleaseMap[name] = bc
}

func BlockChainMapRun(blockChainMsgCh chan<- BlockChainReleaseMsg) {
	for _, bc := range blockChainReleaseMap {
		bc.RunWithChan(blockChainMsgCh)
	}
}
