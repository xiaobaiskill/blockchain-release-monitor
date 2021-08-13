package okexchain

import (
	"github.com/stretchr/testify/assert"
	block_chain "github.com/xiaobaiskill/blockchain-release-monitor/pkg/block-chain-release"
	"testing"
	"time"
)

func TestOkex_RunWithChan(t *testing.T) {
	eth := NewOkex(time.Hour)
	bc := make(chan block_chain.BlockChainReleaseMsg)
	eth.RunWithChan(bc)

	res := <-bc
	t.Log(res)
	assert.Equal(t, res.Name, name)
}
