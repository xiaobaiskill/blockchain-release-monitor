package notification

import "context"

var (
	notices = make([]notificator, 0, 2)
)

func RegisterNotice(n ...notificator) {
	notices = append(notices, n...)
}

func RecvMsg(ctx context.Context, msg Msg) {
	for _, notice := range notices {
		notice.Send(ctx, msg)
	}
}
