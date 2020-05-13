package main

import (
	"context"
	"log"
	"net"

	"github.com/quannv132/practice_0/server/blockchain"

	"github.com/quannv132/practice_0/proto"

	"google.golang.org/grpc"
)

//Server ..
type Server struct {
	Blockchain *blockchain.Blockchain
}

// AddBlock ..
func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

// GetBlockChain ..
func (s *Server) GetBlockChain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash:          b.Hash,
			Data:          b.Data,
		})
	}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Unable to listen on 8080 port:%v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockChain(),
	})
	srv.Serve(listener)
}
