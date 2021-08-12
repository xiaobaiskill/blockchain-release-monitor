package handler

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	api "github.com/xiaobaiskill/blockchain-release-monitor/api/protos/blockchain/v1alpha"
	"github.com/xiaobaiskill/blockchain-release-monitor/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	PublicChainReleasedHandler   *publicChainBlock
	InternalChainReleasedHandler *internalChainBlock
)

func InitGlobal() {
	PublicChainReleasedHandler = NewPublicChainBlock(service.BlockChainSvc)
	InternalChainReleasedHandler = NewInternalChainBlock(service.BlockChainSvc)
}

func RegisterGRPC(s *grpc.Server) {
	api.RegisterInternalReleaseServer(s, InternalChainReleasedHandler)
	api.RegisterPublicReleaseServer(s, PublicChainReleasedHandler)
}

func MustRegisterREST(mux *runtime.ServeMux, grpcListAddress string) {
	ctx := context.Background()
	if err := api.RegisterPublicReleaseHandlerFromEndpoint(ctx, mux, "localhost"+grpcListAddress, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		log.Fatal("register public block-block-released error", zap.Error(err))
	}
}
