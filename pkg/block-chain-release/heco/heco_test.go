package heco

import (
	"github.com/stretchr/testify/assert"
	block_chain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	"testing"
	"time"
)

func TestHeco_RunWithChan(t *testing.T) {
	eth := NewHeco(time.Hour)
	bc := make(chan block_chain.BlockChainReleaseMsg)
	eth.RunWithChan(bc)

	res := <-bc
	t.Log(res)
	assert.Equal(t, res.Name, name)
}
