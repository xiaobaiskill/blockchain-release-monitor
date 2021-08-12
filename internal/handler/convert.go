package handler

import (
	api "github.com/xiaobaiskill/blockchain-release-monitor/api/protos/blockchain/v1alpha"
	"github.com/xiaobaiskill/blockchain-release-monitor/internal/model"
	tsp "google.golang.org/protobuf/types/known/timestamppb"
)

func ReleasesToPB(in []*model.ReleaseInfo) []*api.BlockChain {
	blockChainMap := make(map[string]*api.BlockChain)

	for _, v := range in {
		res, ok := blockChainMap[v.Name]
		if ok {
			res.Releases = append(res.Releases, ReleaseToPB(v))
			blockChainMap[v.Name] = res
		} else {
			blockChainMap[v.Name] = &api.BlockChain{
				Name:     v.Name,
				Releases: []*api.ReleaseInfo{ReleaseToPB(v)},
			}
		}
	}
	out := make([]*api.BlockChain, 0, len(blockChainMap))
	for _, v := range blockChainMap {
		out = append(out, v)
	}
	return out
}

func ReleaseToPB(in *model.ReleaseInfo) *api.ReleaseInfo {
	out := &api.ReleaseInfo{
		ProjectName: in.Project,
		Version:     in.Version,
		Url:         in.Url,
	}

	out.Time = tsp.New(in.Time)

	return out
}
