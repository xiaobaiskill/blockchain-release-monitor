package work_wechat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/notification"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type workWeChat struct {
	url      string
	ch       chan notification.Msg
	retryNum int
	client   *http.Client
}

func RegisterNotice(url string, retryNum int) {
	wx := newWorkWeChat(url, retryNum)
	notification.RegisterNotice(wx)
	go wx.handleMsg()
}

func newWorkWeChat(url string, retryNum int) *workWeChat {
	return &workWeChat{
		url:      url,
		ch:       make(chan notification.Msg, 10),
		retryNum: retryNum,
		client: &http.Client{
			Timeout: time.Second * 3,
		},
	}
}

func (w *workWeChat) handleMsg() {
	var err error
	for msg := range w.ch {
		for i := 0; i < w.retryNum; i++ {
			if err = w.send(msg); err == nil {
				i = w.retryNum
			} else {
				log.Error("failed to send err", zap.Error(err))
				i++
			}
		}
	}
}

func (w *workWeChat) Send(_ context.Context, msg notification.Msg) {
	w.ch <- msg
}

func (w *workWeChat) send(msg notification.Msg) error {
	msgBytes, err := w.generateSendMsg(msg.Title, msg.Level, msg.Data)
	if err != nil {
		return fmt.Errorf("failed to generateSendMsg err:%v", err)
	}

	resp, err := w.client.Post(w.url, "application/json", bytes.NewReader(msgBytes))
	if err != nil {
		return fmt.Errorf("failed to post request err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("wrong response: %d", resp.StatusCode)
	}
	return nil
}

func (w *workWeChat) generateSendMsg(subject string, level notification.Level, noticeData map[string]interface{}) ([]byte, error) {
	color := w.getLevel(level)
	str := fmt.Sprintf("**<font color='%s'>%v</font>**\n", color, subject)
	if len(noticeData) > 0 {
		for k, v := range noticeData {
			str += fmt.Sprintf("> %s: <font color='%s'>%v</font> \n", k, color, v)
		}
	}
	data := &wxData{
		Msgtype: "markdown",
		Markdown: markDwon{
			Content: str,
		},
	}
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return byteData, nil
}

func (*workWeChat) getLevel(level notification.Level) string {
	color := "comment"
	switch level {
	case notification.Info:
		color = "info"
	case notification.Warning, notification.Error:
		color = "warning"
	}
	return color
}

type wxData struct {
	Markdown markDwon `json:"markdown"`
	Msgtype  string   `json:"msgtype"`
}

type markDwon struct {
	Content string `json:"content"`
}
