package etherum

import (
	"github.com/stretchr/testify/assert"
	block_chain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	"testing"
)

func TestEth_RunWithChan(t *testing.T) {
	eth := newEth()
	bc := make(chan block_chain.BlockChainReleaseMsg)
	eth.RunWithChan(bc)

	res := <-bc
	assert.Equal(t, res.Name, name)
}
