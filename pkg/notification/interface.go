package notification

import "context"

type Level int32

const (
	Info Level = iota
	Warning
	Error
)

type Msg struct {
	Title string
	Level Level
	Data  map[string]interface{}
}

type notificator interface {
	Send(context.Context, Msg)
}
