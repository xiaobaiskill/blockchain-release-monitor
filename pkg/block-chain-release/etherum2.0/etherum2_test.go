package etherum2

import (
	"github.com/stretchr/testify/assert"
	block_chain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	"testing"
)

func TestEth2_RunWithChan(t *testing.T) {
	eth := newEth2()
	bc := make(chan block_chain.BlockChainReleaseMsg)
	eth.RunWithChan(bc)

	res := <-bc
	t.Log(res)
	assert.Equal(t, res.Name, name)
}
