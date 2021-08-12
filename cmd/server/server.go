package server

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/xiaobaiskill/blockchain-release-monitor/internal/handler"
	"github.com/xiaobaiskill/blockchain-release-monitor/internal/service"
	"github.com/xiaobaiskill/blockchain-release-monitor/pkg/notification/slack"
	work_wechat "github.com/xiaobaiskill/blockchain-release-monitor/pkg/notification/work-wechat"
	"github.com/xiaobaiskill/kit/app"
	"github.com/xiaobaiskill/kit/rest"
	"github.com/xiaobaiskill/kit/rpc"
	"go.uber.org/zap"
	"io/ioutil"
)

func NewServer() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "server",
		Short: "",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		Run:   run,
	}
	return cmd
}

func initConfig(file string, cfg *serverConfig) {
	fileByte, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("invalid file")
	}
	err = json.Unmarshal(fileByte, cfg)
	if err != nil {
		log.Fatal("failed to json unmarshal", zap.Error(err))
	}
}

func run(_ *cobra.Command, args []string) {
	var cfg serverConfig
	initConfig(args[0], &cfg)

	if cfg.Notification.WorkWechat != nil {
		work_wechat.RegisterNotice(*cfg.Notification.WorkWechat)
	}
	if cfg.Notification.Slack != nil {
		slack.RegisterNotice(*cfg.Notification.Slack)
	}

	log.Info("init service")
	service.GlobalInit()

	log.Info("init handler")
	handler.InitGlobal()

	log.Info("register grpc")
	rpcServer := rpc.NewServer(&rpc.Config{ListenAddress: cfg.GrpcAddr})
	handler.RegisterGRPC(rpcServer.Server)

	log.Info("grpc runnig...")
	rpcServer.MustListenAndServe()

	log.Info("register rest")
	restServer := rest.NewServer(&rest.Config{
		CORSMaxAge:    86400,
		ListenAddress: cfg.RestAddr,
		EmitDefaults:  true,
	})
	handler.MustRegisterREST(restServer.ServeMux, rpcServer.Address)

	restServer.ListenAndServed()

	app.Exit()
}
