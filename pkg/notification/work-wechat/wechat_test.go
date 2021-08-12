package work_wechat

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/notification"
	"testing"
)

func TestWorkWeChat_Send(t *testing.T) {
	testwx := newWorkWeChat("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=XXXX", 3)
	err := testwx.send(notification.Msg{
		Title: "test",
		Level: notification.Error,
		Data: map[string]interface{}{
			"test1": "test1",
			"test2": "test2",
		},
	})
	assert.NoError(t, err)
}
