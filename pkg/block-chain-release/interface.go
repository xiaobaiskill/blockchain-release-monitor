package block_chain_release

type BlockChainReleaseMsg struct {
	Name    string
	Project string
	Version string
	Url     string
}

type blockChainReleaser interface {
	RunWithChan(chan<- BlockChainReleaseMsg)
}
