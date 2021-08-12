package server

type Notification struct {
	WorkWechat *string `json:"work_wechat"`
	Slack      *string `json:"slack"`
}

type serverConfig struct {
	RestAddr     string        `json:"rest_listen"`
	GrpcAddr     string        `json:"grpc_listen"`
	Notification *Notification `json:"notification"`
}
