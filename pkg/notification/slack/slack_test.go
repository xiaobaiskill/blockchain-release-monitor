package slack

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/notification"
	"testing"
)

func TestWorkWeChat_Send(t *testing.T) {
	testwx := newSlack("https://hooks.slack.com/services/XXXX", 3)
	err := testwx.send(notification.Msg{
		Title: "binance smart chain",
		Level: notification.Info,
		Data: map[string]interface{}{
			"project": "binance-chain",
			"version": "v1.1.1",
			"url":     "https://github.com/binance-chain/bsc/releases/latest",
		},
	})
	assert.NoError(t, err)
}
