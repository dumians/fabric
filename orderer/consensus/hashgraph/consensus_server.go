/*
SPDX-License-Identifier: Apache-2.0
*/

package hashgraph

import (
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/protos/orderer"
	ctx "golang.org/x/net/context"
)

const (
	pkgLogID = "orderer/consensus/hashgraph"
)

var (
	logger = flogging.MustGetLogger(pkgLogID)
)

type ordererFeedServer struct{}

func New() orderer.OrdererServiceServer {
	return &ordererFeedServer{}
}

func (s *ordererFeedServer) Consensus(ctx ctx.Context, in *orderer.ConsensusTransaction) (*orderer.ConsensusResponse, error) {
	logger.Info("HELLO From Hashgraph!")
	return &orderer.ConsensusResponse{}, nil
}
