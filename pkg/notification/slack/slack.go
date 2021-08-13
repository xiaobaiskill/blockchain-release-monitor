package slack

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

type slack struct {
	url      string
	ch       chan notification.Msg
	retryNum int
	client   *http.Client
}

func RegisterNotice(url string, retryNum int) {
	s := newSlack(url, retryNum)
	notification.RegisterNotice(s)
	go s.handleMsg()
}

func newSlack(url string, retryNum int) *slack {
	return &slack{
		url:      url,
		ch:       make(chan notification.Msg, 10),
		retryNum: retryNum,
		client: &http.Client{
			Timeout: time.Second * 3,
		},
	}
}
func (s *slack) handleMsg() {
	var err error
	for msg := range s.ch {
		for i := 0; i < s.retryNum; i++ {
			if err = s.send(msg); err == nil {
				i = s.retryNum
			} else {
				log.Error("failed to send err", zap.Error(err))
				i++
			}
		}
	}
}

func (s *slack) Send(_ context.Context, msg notification.Msg) {
	s.ch <- msg
}

func (s *slack) send(msg notification.Msg) error {
	msgBytes, err := s.generateSendMsg(msg.Title, msg.Level, msg.Data)
	if err != nil {
		return fmt.Errorf("failed to generateSendMsg err:%v", err)
	}

	resp, err := s.client.Post(s.url, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		return fmt.Errorf("failed to post request err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("wrong response: %d", resp.StatusCode)
	}
	return nil
}

func (s *slack) generateSendMsg(subject string, level notification.Level, noticeData map[string]interface{}) ([]byte, error) {
	blocks := make([]slackNoticeBlock, 0, len(noticeData)+1)
	blocks = append(blocks, slackNoticeBlock{
		Text: slackNoticeText{
			Emoji: true,
			Text:  subject,
			Type:  "plain_text",
		},
		Type: "header",
	})

	if len(noticeData) != 0 {
		str := ""
		for k, v := range noticeData {
			str += fmt.Sprintf("*%s:* %v\n", k, v)
		}

		blocks = append(blocks, slackNoticeBlock{
			Text: slackNoticeText{
				Text: str,
				Type: "mrkdwn",
			},
			Type: "section",
		})
	}

	data := slackNotice{
		Attachments: []slackNoticeAttachment{
			{
				Blocks: blocks,
				Color:  s.getLevel(level),
			},
		},
	}
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return byteData, nil
}

func (s *slack) getLevel(level notification.Level) string {
	var color string
	switch level {
	case notification.Info:
		color = "#86ff85"
	case notification.Warning:
		color = "#fc4b40"
	case notification.Error:
		color = "#f52727"
	default:
		color = "#c1bebd"
	}
	return color
}

type slackNotice struct {
	Attachments []slackNoticeAttachment `json:"attachments"`
}

type slackNoticeAttachment struct {
	Blocks []slackNoticeBlock `json:"blocks"`
	Color  string             `json:"color"`
}

type slackNoticeBlock struct {
	Text slackNoticeText `json:"text"`
	Type string          `json:"type"`
}

type slackNoticeText struct {
	Emoji bool   `json:"emoji,omitempty"`
	Text  string `json:"text"`
	Type  string `json:"type"`
}
