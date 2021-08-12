package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/xiaobaiskill/blockchain-release-monitor/api/protos/blockchain/v1alpha"
	"github.com/xiaobaiskill/blockchain-release-monitor/internal/service"
)

type publicChainBlock struct {
	svc service.BlockChainer
	api.UnimplementedPublicReleaseServer
}

func NewPublicChainBlock(svc service.BlockChainer) *publicChainBlock {
	return &publicChainBlock{
		svc: svc,
	}
}

func (i *publicChainBlock) GetAllRelease(ctx context.Context, _ *empty.Empty) (*api.GetAllData, error) {
	data := i.svc.GetReleaseData(ctx)
	return &api.GetAllData{Data: ReleasesToPB(data)}, nil
}
