package register

import (
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/notification/slack"
	work_wechat "github.com/xiaobaiskill/blockchain-release-monitor/pkg/notification/work-wechat"
)

type Webhooks struct {
	Slack  string
	WeChat string
}

const retryNum = 3

func Register(w Webhooks) {
	if w.Slack != "" {
		log.Info("register slack")
		slack.RegisterNotice(w.Slack, retryNum)
	}
	if w.WeChat != "" {
		log.Info("register wechat")
		work_wechat.RegisterNotice(w.WeChat, retryNum)
	}
}
