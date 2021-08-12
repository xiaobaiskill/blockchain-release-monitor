package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/xiaobaiskill/blockchain-release-monitor/api/protos/blockchain/v1alpha"
	"github.com/xiaobaiskill/blockchain-release-monitor/internal/service"
)

type internalChainBlock struct {
	svc service.BlockChainer
	api.UnimplementedInternalReleaseServer
}

func NewInternalChainBlock(svc service.BlockChainer) *internalChainBlock {
	return &internalChainBlock{
		svc: svc,
	}
}

func (i *internalChainBlock) GetAllRelease(ctx context.Context, _ *empty.Empty) (*api.GetAllData, error) {
	data := i.svc.GetReleaseData(ctx)
	return &api.GetAllData{Data: ReleasesToPB(data)}, nil
}
